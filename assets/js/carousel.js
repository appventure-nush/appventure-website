// (c) 2016 Ambrose Chua and Wayne Tee
// Released for free under the WTFPL

function Carousel(ele) {
  if (typeof ele == "string") {
    var carousels = document.querySelectorAll(ele);
    if (carousels.length > 1) {
      var objs = [];
      for (var i = 0; i < carousels.length; i++) {
        objs.push(new Carousel(carousels[i]));
      }
      return objs;
    } else if (carousels.length == 1) {
      return new Carousel(carousels[0]);
    }
    return;
  }

  this.parent = ele;
  this.items = ele.querySelectorAll(".carousel-item");
  this.triggers = ele.querySelectorAll("[data-index]");
  this.index = ele.dataset.start * 1 || 0;
  this.autoplay.disableOnClick = true;
  this.autoplay.enabled = false;
  this.autoplay.timer = 0;
  this.autoplay.durationDefault = 3000;
  this.autoplay.duration = this.autoplay.durationDefault;
  this._registerEvents();
  this._registerTouchEvents();
  this.change(this.index);
  if (this.items.length < 2) {
    this.parent.classList.add("one-item");
  }
}

Carousel.prototype.change = function(id, isClick) {
  var idIsString = typeof id == "string";
  if (idIsString && id.charAt(0) == "+") {
    return this.change(this.index + parseInt(id.substring(1)), isClick);
  } else if (idIsString && id.charAt(0) == "-") {
    return this.change(this.index - parseInt(id.substring(1)), isClick);
  }

  this.index = parseInt(id);
  if (this.index > this.items.length - 1) {
    this.index = 0;
  } else if (this.index < 0) {
    this.index = this.items.length - 1;
  }

  for (var i = 0; i < this.items.length; i++) {
    this.parent.classList.remove(
      "carousel-type-" + this.items[i].dataset["type"]
    );
  }

  for (var i = 0; i < this.items.length; i++) {
    this.items[i].classList.remove("active");
    if (this.index == i) {
      this.items[i].classList.add("active");
      this.parent.title = this.items[i].children[0].title;
      this.parent.classList.add(
        "carousel-type-" + this.items[i].dataset["type"]
      );
    }
  }

  for (var i = 0; i < this.triggers.length; i++) {
    this.triggers[i].classList.remove("active");
    if (this.index == this.triggers[i].dataset.index) {
      this.triggers[i].classList.add("active");
    }
  }

  if (isClick) {
    if (this.autoplay.disableOnClick) {
      this.autoplay(-1);
    }
  }

  if (this.autoplay.enabled) {
    this.autoplay();
  }
};

Carousel.prototype._registerEvents = function() {
  var that = this;
  for (var i = 0; i < this.triggers.length; i++) {
    this.triggers[i].addEventListener("click", function(e) {
      that.change(this.dataset.index, true);
      e.preventDefault();
      e.stopPropagation();
    });
  }
  this.parent.querySelector(".carousel-items").addEventListener(
    "click",
    function(e) {
      window.open(that.items[that.index].dataset["href"]);
      e.preventDefault();
    },
    false
  );
};

Carousel.prototype._registerTouchEvents = function() {
  var that = this;

  var touches = {
    start: { x: 0, y: 0 },
    end: { x: 0, y: 0 },
    hasmoved: false,
  };
  this.parent.addEventListener("touchstart", function(e) {
    touches.start.x = e.touches[0].clientX;
    touches.start.y = e.touches[0].clientY;
  });
  this.parent.addEventListener("touchmove", function(e) {
    touches.end.x = e.touches[0].clientX;
    touches.end.y = e.touches[0].clientY;
    touches.hasmoved = true;
    e.preventDefault();
  });
  this.parent.addEventListener("touchend", function() {
    var diff = {
      x: touches.end.x - touches.start.x,
      y: touches.end.y - touches.start.y,
    };
    if (touches.hasmoved && Math.abs(diff.x) > Math.abs(diff.y)) {
      if (diff.x > 64) {
        that.change("-1", true);
      } else if (diff.x < -64) {
        that.change("+1", true);
      }
    }
    touches.hasmoved = false;
  });
};

Carousel.prototype.autoplay = function(duration) {
  var that = this;

  this.autoplay.duration = duration || this.autoplay.durationDefault;
  this.autoplay.enabled = this.autoplay.duration > 0 ? true : false;

  clearTimeout(this.autoplay.timer);
  if (this.autoplay.enabled) {
    this.autoplay.timer = setTimeout(function() {
      that.change("+1");
    }, this.autoplay.duration);
  }
};
