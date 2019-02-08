function ScrollManager() {
  this.onScroll = this.onScroll.bind(this);
  this.onScroll();
}

ScrollManager.prototype.onScroll = function(e) {
  var elements = this._getElements();
  for (var i = 0; i < elements.length; i++) {
    var element = elements[i];
    if (this._inViewport(element)) {
      if (!element.classList.contains("start")) {
        var newelement = element.cloneNode(true); // force rerender, Safari bug.
        newelement.classList.add("start");
        element.parentNode.replaceChild(newelement, element);
      }
    }
  }
};

ScrollManager.prototype._getElements = function() {
  return document.querySelectorAll(".animate");
};

ScrollManager.prototype._inViewport = function(ele) {
  var scrollPosition =
    document.documentElement.scrollTop + document.documentElement.clientHeight;
  var elePosition = ele.offsetTop;
  var buffer = 64;
  return elePosition < scrollPosition - buffer;
};
