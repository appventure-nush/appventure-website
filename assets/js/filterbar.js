var opened = false;

function FilterBar(ele, collection) {
  if (typeof ele == "string") {
    var filterbars = document.querySelectorAll(ele);
    if (filterbars.length > 1) {
      var objs = [];
      for (var i = 0; i < filterbars.length; i++) {
        objs.push(new FilterBar(filterbars[i], collection));
      }
      return objs;
    } else if (filterbars.length == 1) {
      return new FilterBar(filterbars[0], collection);
    }
    return;
  }

  this.parent = ele;
  this._registerEvents();
  this.filter = new Filter(collection);
}

FilterBar.prototype._registerEvents = function() {
  var that = this;

  var groups = this.parent.querySelectorAll(".group");
  for (var i = 0; i < groups.length; i++) {
    groups[i].querySelector("label").addEventListener("click", function(e) {
      e.stopPropagation();
      for (var j = 0; j < groups.length; j++) {
        if (this.parentElement != groups[j]) {
          groups[j].classList.remove("opened");
        }
      }
      this.parentElement.classList.toggle("opened");
    });
    var tags = groups[i].querySelectorAll("[data-tag]");
    for (var j = 0; j < tags.length; j++) {
      tags[j].addEventListener("click", function(e) {
        e.stopPropagation();
        this.classList.toggle("checked");
        that._updateFilter();
      });
    }
  }

  document.body.addEventListener("click", function() {
    for (var j = 0; j < groups.length; j++) {
      groups[j].classList.remove("opened");
    }
  });
};

FilterBar.prototype._updateFilter = function() {
  var filter = [];
  var groups = this.parent.querySelectorAll(".group");
  for (var i = 0; i < groups.length; i++) {
    var groupfilter = [];
    var tags = groups[i].querySelectorAll("[data-tag].checked");
    for (var j = 0; j < tags.length; j++) {
      groupfilter.push(tags[j].dataset.tag);
    }
    if (groupfilter.length > 0) {
      filter.push(groupfilter);
    }
  }
  console.log(filter);
  this.filter.update(filter);
};

document.onclick = function(){
  if(opened){
    var filterbar = document.querySelector("div.filterbar");
    var groups = filterbar.querySelectorAll(".group");
    for (var j = 0; j < groups.length; j++) {
      if(groups[j].classList.contains("opened")){
        groups[j].classList.remove("opened");
        opened = false;
      }
    }
  }else{
    opened = true;
  }
};
