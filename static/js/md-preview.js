// Depends: markdown.js

function Editor(input, preview) {
	this.update = function () {
		/*
		preview.innerHTML = markdown.toHTML(input.value);
		*/
		preview.innerHTML = marked(input.value);
		input.style.height = 'auto';
		input.style.height = input.scrollHeight + "px";
		preview.style.height = input.style.height;
	};
	input.editor = this;
	this.update();
}

var $ = function (id) {
	return document.getElementById(id);
};

