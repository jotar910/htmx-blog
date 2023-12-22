class PostComponent extends HTMLElement {

    connectedCallback(): void {
        (this.querySelectorAll('img') as NodeListOf<HTMLImageElement>).forEach((el) => {
            let parent: HTMLElement | null = this;
            while (parent !== null && !this.hasAttribute('data-post')) {
                parent = parent.parentElement;
            }
            if (parent) {
                el.src = `articles/${this.getAttribute('data-post')}/${el.src.split('/').at(-1)}`;
            }
            el.loading = 'lazy';
        });
    }

}

customElements.define('blog-post', PostComponent);
