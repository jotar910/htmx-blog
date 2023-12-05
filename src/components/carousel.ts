const AUTO_CHANGE_PAGE_INTERVAL: number = 5000;

class CarouselComponent extends HTMLElement {

    shadow!: ShadowRoot;
    wrapper!: HTMLElement;
    indicators!: HTMLElement;
    leftArrow!: HTMLElement;
    rightArrow!: HTMLElement;

    totalItems!: number;
    itemsPerPage!: number;

    canScroll: boolean = false;
    totalPages!: number;
    nextPageTimer!: NodeJS.Timeout;

    get currentItemIndex(): number {
        return this._currentItemIndex;
    }
    set currentItemIndex(v: number) {
        this._currentItemIndex = v;
        this.updateIndicatorActive();
    }
    _currentItemIndex = 0;


    private set canHideAnimation(canHideAnimation: boolean) {
        if (canHideAnimation) {
            this.wrapper.classList.add('hide-animation');
        } else {
            this.wrapper.classList.remove('hide-animation');
        }
    }

    constructor() {
        super();
        this.shadow = this.attachShadow({ mode: "open" });
        this.shadow.innerHTML = `
        <style>
            @import "dist/main.css";
        </style>
        <template id="container">
            <div class="gallery-list-wrapper">
              <slot name="item"></slot>
              <a class="scroll-button left-scroll-button">
                <img src="assets/icons/left-arrow.svg" alt="left" class="scroll-arrow" />
              </a>
              <a class="scroll-button right-scroll-button">
                <img src="assets/icons/right-arrow.svg" alt="right" class="scroll-arrow" />
              </a>
              <div class="pages-indicator"></div>
            </div>
        </template>`;
        const containerTemplate = (this.shadow.getElementById(
            "container",
        ) as HTMLTemplateElement).content;
        this.shadow.appendChild(containerTemplate.cloneNode(true));
    }

    connectedCallback(): void {
        this.totalItems = this.children.length;
        this.itemsPerPage = 1;

        this.wrapper = this.shadow.querySelector('.gallery-list-wrapper') as HTMLElement;
        this.indicators = this.shadow.querySelector('.pages-indicator') as HTMLElement;
        this.leftArrow = this.shadow.querySelector('.left-scroll-button > .scroll-arrow') as HTMLElement;
        this.rightArrow = this.shadow.querySelector('.right-scroll-button > .scroll-arrow') as HTMLElement;

        this.wrapper.addEventListener('mouseenter', this.onWrapperMouseEnter.bind(this));
        this.leftArrow.addEventListener('click', this.onScrollLeft.bind(this));
        this.rightArrow.addEventListener('click', this.onScrollRight.bind(this));

        this.updateTotalPages();
        this.updateItemWidthCssVariable();
        this.updateTotalPages();
        this.appendIndicatorElements();

        this.canScroll = this.totalItems > 1;
        if (this.canScroll) {
            this.leftArrow.classList.remove('hidden');
            this.rightArrow.classList.remove('hidden');
            this.setNextPageTimer(); // First page index is 0.
        } else {
            this.leftArrow.classList.add('hidden');
            this.rightArrow.classList.add('hidden');
        }
        this.updatePageCssVariable();
        this.updateItemWidthCssVariable();
    }

    disconnectedCallback(): void {
        for (let i = 0; i < this.indicators.children.length; ++i) {
            this.indicators.children[i].removeEventListener('click', this.changePageHandler(i).bind(this));
        }
        this.rightArrow.addEventListener('click', this.onScrollRight.bind(this));
        this.leftArrow.addEventListener('click', this.onScrollLeft.bind(this));
        this.wrapper.addEventListener('mouseenter', this.onWrapperMouseEnter.bind(this));
        clearTimeout(this.nextPageTimer);
    }

    onScrollLeft(): void {
        this.updateStateOnChangePage('left');
        this.updatePageCssVariable();
        this.setNextPageTimer();
    }

    onScrollRight(): void {
        this.updateStateOnChangePage('right');
        this.updatePageCssVariable();
        this.setNextPageTimer();
    }

    changePage(pageIndex: number): void {
        this.currentItemIndex = pageIndex;
        this.updatePageCssVariable();
        this.setNextPageTimer();
    }

    setNextPageTimer(): void {
        clearTimeout(this.nextPageTimer);
        this.nextPageTimer = setTimeout(() =>
            this.changePage((this.currentItemIndex + 1) % this.totalPages),
            AUTO_CHANGE_PAGE_INTERVAL);
    }

    onWrapperMouseEnter() {
        this.canHideAnimation = true;
    }


    private updateItemWidthCssVariable(): void {
        this.style.setProperty('--item-width', `${100 / (this.itemsPerPage || 1)}%`);
    }

    private updatePageCssVariable(): void {
        this.style.setProperty('--horizontal-page', `${this.currentItemIndex}`);
    }

    private updateStateOnChangePage(side: 'left' | 'right'): void {
        this.currentItemIndex = (this.itemsPerPage * (side === 'right' ? 1 : -1) + this.currentItemIndex + this.totalItems) % this.totalItems;
    }

    private updateTotalPages(): void {
        this.totalPages = Math.ceil(this.totalItems / this.itemsPerPage);
    }

    private updateIndicatorActive(): void {
        for (let i = 0; i < this.indicators.children.length; ++i) {
            if (i === this.currentItemIndex) {
                this.indicators.children[i].classList.add('active');
            } else {
                this.indicators.children[i].classList.remove('active');
            }
        }
    }

    private appendIndicatorElements() {
        Array(this.totalPages).fill(0).forEach((_, i) => {
            const el = document.createElement('div');
            el.className = 'page-indicator';
            i === 0 && el.classList.add('active');
            el.addEventListener('click', this.changePageHandler(i).bind(this));
            this.indicators.appendChild(el);
        });
    }

    private changePageHandler(index: number): () => void {
        const cache: (() => void)[] = [];
        if (!cache[index]) {
            cache[index] = () => this.changePage(index);
        }
        return cache[index];
    }

}

if (customElements.get("app-carousel") === undefined) {
    customElements.define("app-carousel", CarouselComponent);
}
