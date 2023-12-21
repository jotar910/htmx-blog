class ImageComponent extends HTMLImageElement {
    connectedCallback() {
        let parent = this.parentElement;
        while (parent !== null && !parent.hasAttribute('data-post')) {
            parent = parent.parentElement;
        }
        if (!parent) {
            return;
        }
        this.src = `articles/${parent.getAttribute('data-post')}/${this.src.split('/').at(-1)}`;
    }
}

customElements.define("blog-image", ImageComponent, { extends: "img" });
