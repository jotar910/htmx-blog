class PinterestLayoutComponent extends HTMLElement {

    observer!: MutationObserver;
    previousColCount = 0;
    previousTotalHeight = 0;

    connectedCallback() {
        this.observeChildren();
        window.addEventListener('resize', this.updateLayout.bind(this));
        this.updateLayout();
    }

    disconnectedCallback(): void {
        window.removeEventListener('resize', this.updateLayout.bind(this));
        this.observer.disconnect();
    }

    private observeChildren(): void {
        this.observer = new MutationObserver(() => {
            this.previousColCount = 0;
            this.previousTotalHeight = 0;
            this.updateLayout();
        });

        this.observer.observe(this, { childList: true });
    }

    private updateLayout() {
        const COL_MARGIN = 30;
        const COL_COUNT = this.calculateColCount(this.clientWidth, 360);
        const COL_WIDTH = this.calculateColWidth(this.clientWidth, COL_COUNT);

        if (COL_COUNT === this.previousColCount) {
            let totalHeight = 0;
            for (let i = 0; i < this.children.length; i++) {
                const child = this.children[i] as HTMLElement;
                child.style.maxWidth = `${COL_WIDTH}px`;
                totalHeight += child.offsetHeight;
            }
            if (totalHeight === this.previousTotalHeight) {
                return;
            }
            this.previousTotalHeight = totalHeight;
        }

        this.previousColCount = COL_COUNT;

        const colHeights = new Array(COL_COUNT).fill(0);
        const colLastEl: HTMLElement[] = new Array(COL_COUNT).fill(null);
        let totalHeight = 0;

        let order = 0;
        for (let col = 0; col < COL_COUNT; ++col) {
            for (let i = col; i < this.children.length; i += COL_COUNT) {
                const child = this.children[i] as HTMLElement;
                child.style.order = `${order}`;
                child.style.marginBottom = `${COL_MARGIN}px`;
                child.style.maxWidth = `${COL_WIDTH}px`;
                colHeights[col] += child.offsetHeight + COL_MARGIN;
                colLastEl[col] = child;
                totalHeight += child.offsetHeight;
                order++;
            }
        }

        this.previousTotalHeight = totalHeight;

        const highest = Math.max.apply(Math, colHeights);
        this.style.height = highest + 'px';

        for (let i = 0; i < colLastEl.length; ++i) {
            if (colLastEl[i]) {
                colLastEl[i].style.marginBottom = `${highest - colHeights[i]}px`;
            } else {
                const placeholderEl = document.createElement('article');
                placeholderEl.style.width = '100%';
                this.appendChild(placeholderEl);
            }
        }
    }

    private calculateColCount(containerWidth: number, minItemWidth: number): number {
        const COL_MARGIN = 30;
        let cols = 0;
        let width = -COL_MARGIN;
        while (width < containerWidth) {
            width += minItemWidth + COL_MARGIN;
            cols++;
        }
        return Math.max(1, cols - 1);
    }

    private calculateColWidth(containerWidth: number, colCount: number): number {
        const COL_MARGIN = 30;
        return Math.max(360, Math.floor((containerWidth - (colCount - 1) * COL_MARGIN) / colCount));
    }
}

customElements.define("app-pinterest-layout", PinterestLayoutComponent);
