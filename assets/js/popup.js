function Popup(hash, content, onshow, onhide) {
  this.hash = hash;
  this.content = content;
  this._ele = null;
  this.onshow = onshow;
  this.onhide = onhide;

  this._hashChange = this._hashChange.bind(this);
  window.addEventListener("hashchange", this._hashChange);
  this._hashChange();
}

Popup.prototype._hashChange = function() {
  if (window.location.hash.replace("#", "") == this.hash) {
    this.show();
  } else {
    this.hide();
  }
};

Popup.prototype.show = function() {
  if (this._ele) {
    this.hide();
  }

  var popupWrapper = document.createElement("div");
  popupWrapper.classList.add("popup-wrapper");
  var popup = document.createElement("div");
  popup.classList.add("popup");
  popup.innerHTML = this.content;
  var closeBtnWrapper = document.createElement("div");
  closeBtnWrapper.classList.add("close");
  var closeBtn = document.createElement("a");
  closeBtn.classList.add("icon-close");
  closeBtn.href = "#";

  closeBtnWrapper.appendChild(closeBtn);
  popupWrapper.appendChild(closeBtnWrapper);
  popupWrapper.appendChild(popup);
  document.body.appendChild(popupWrapper);
  this._ele = popupWrapper;

  if (typeof this.onshow == "function") {
    this.onshow();
  }
};

Popup.prototype.hide = function() {
  if (!this._ele) {
    return;
  }

  document.body.removeChild(this._ele);
  this._ele = null;

  if (typeof this.onhide == "function") {
    this.onhide();
  }
};
