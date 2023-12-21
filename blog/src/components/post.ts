class PostComponent extends HTMLDivElement {

    connectedCallback(): void {
        (this.querySelectorAll('img') as NodeListOf<HTMLImageElement>).forEach((el) => {
            const image = document.createElement('img', { is: 'blog-image' });
            image.src = el.src;
            for (let attr of el.attributes) {
                image.setAttribute(attr.name, attr.value);
            }
            el.replaceWith(image);
        });
    }

}

customElements.define("blog-post", PostComponent, { extends: "div" });
