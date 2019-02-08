function Filter(ele) {
  if (typeof ele == "string") {
    var filters = document.querySelectorAll(ele);
    if (filters.length > 1) {
      var objs = [];
      for (var i = 0; i < filters.length; i++) {
        objs.push(new Filter(filters[i]));
      }
      return objs;
    } else if (filters.length == 1) {
      return new Filter(filters[0]);
    }
    return;
  }

  this.parent = ele;
  this.items = ele.querySelectorAll("[data-filterable]");
}

Filter.prototype.update = function(tags) {
  var matches = this._matchTags(tags);
  this._updateElements(matches, this.items);
};

Filter.prototype._matchTags = function(tags) {
  var matches = [];
  for (var i = 0; i < this.items.length; i++) {
    var itemtags = this._resolveTags(this.items[i]);
    if (this._matches(tags, itemtags)) {
      matches.push(this.items[i]);
    }
  }
  return matches;
};

Filter.prototype._matches = function(rule, target) {
  for (var i = 0; i < rule.length; i++) {
    for (var j = 0; j < rule[i].length; j++) {
      if (target.indexOf(rule[i][j]) >= 0) {
        break;
      }
    }
    if (j == rule[i].length) {
      return false;
    }
  }
  return true;
};

Filter.prototype._resolveTags = function(ele) {
  var dataset = ele.dataset;
  var fields = dataset.filterable.split(",");
  var tags = [];
  for (var i = 0; i < fields.length; i++) {
    var field = dataset[fields[i]];
    if (field) {
      field = field.split(",");
      tags = tags.concat(field);
    }
  }
  return tags;
};

Filter.prototype._updateElements = function(matched, all) {
  for (var i = 0; i < all.length; i++) {
    all[i].classList.remove("matched");
    all[i].classList.add("unmatched");
  }
  for (var j = 0; j < matched.length; j++) {
    matched[j].classList.remove("unmatched");
    matched[j].classList.add("matched");
  }
};
