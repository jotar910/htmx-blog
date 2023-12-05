class NetflixHorizontalListComponent extends HTMLElement {

    shadow!: ShadowRoot;
    leftArrowEl!: HTMLElement;
    rightArrowEl!: HTMLElement;

    get items(): HTMLElement[] {
        return (this.children[0]?.children as unknown as HTMLElement[]) || [];
    }

    constructor() {
        super();
        this.shadow = this.attachShadow({ mode: "open" });
        this.shadow.innerHTML = `
        <style>
            @import "dist/main.css";
        </style>
        <template id="content">
            <div class="horizontal-list flex items-center relative">
                <div class="scroll-button z-10 w-12 h-12 left-0 translate-x-[-50%] translate-y-[-60%] bg-base-300 border border-primary-100 shadow-xl rounded-full text-gray-700 dark:text-gray-200 previous">
                    <img src="assets/icons/left-arrow.svg" alt="left" class="scroll-arrow" />
                </div>
                <slot name="list"></slot>
                <div class="scroll-button z-10 w-12 h-12 right-0 translate-x-[50%] translate-y-[-60%] bg-base-300 border border-primary-100 shadow-xl rounded-full text-gray-700 dark:text-gray-200 next">
                     <img src="assets/icons/right-arrow.svg" alt="right" class="scroll-arrow" />
                </div>
            </div>
        </template>`;
        const template = (this.shadow.getElementById(
            "content",
        ) as HTMLTemplateElement).content;
        this.shadow.appendChild(template.cloneNode(true));
    }


    connectedCallback(): void {
        this.leftArrowEl = this.shadow.querySelector('.previous') as HTMLElement;
        this.rightArrowEl = this.shadow.querySelector('.next') as HTMLElement;

        this.leftArrowEl.addEventListener('click', this.onPrevious.bind(this));
        this.rightArrowEl.addEventListener('click', this.onNext.bind(this));

        this.updateArrows();
        window.addEventListener('resize', this.updateArrows.bind(this));
        this.children[0].addEventListener('scroll', this.updateArrows.bind(this));
    }

    disconnectedCallback(): void {
        this.removeEventListener('scroll', this.updateArrows.bind(this));
        this.children[0].removeEventListener('resize', this.updateArrows.bind(this));

        this.rightArrowEl.removeEventListener('click', this.onNext.bind(this));
        this.leftArrowEl.removeEventListener('click', this.onPrevious.bind(this));
    }

    onNext(): void {
        if (this.items.length === 0) {
            return;
        }

        const firstVisible = this.firstVisibleChild();
        if (firstVisible === null) {
            this.items[0].scrollIntoView();
            return;
        }
        const lastVisible = this.lastVisibleChildFrom(firstVisible);
        if (lastVisible === this.items.length - 1) {
            return;
        }
        const perPage = lastVisible - firstVisible + 1;
        this.items[lastVisible + 1].scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'start' });
        this.updateArrowsVisibility(true, lastVisible + perPage < this.items.length);
    }

    onPrevious(): void {
        if (this.items.length === 0) {
            return;
        }

        const firstVisible = this.firstVisibleChild();
        if (firstVisible === null) {
            this.items[0].scrollIntoView();
            return;
        }
        if (firstVisible === 0) {
            return;
        }
        const lastVisible = this.lastVisibleChildFrom(firstVisible);
        const perPage = lastVisible - firstVisible + 1;
        const index = Math.max(0, firstVisible - perPage);
        this.items[index].scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'start' });
        this.updateArrowsVisibility(index !== 0, true);
    }

    private firstVisibleChild(): number | null {
        const containerRect = this.getBoundingClientRect();
        for (let i = 0; i < this.items.length; ++i) {
            const childEl = this.items[i] as HTMLElement;

            if (this.isInViewport(containerRect, childEl.getBoundingClientRect())) {
                return i;
            }
        }
        return null;
    }

    private lastVisibleChildFrom(fromIndex: number): number {
        const containerRect = this.getBoundingClientRect();
        for (let i = fromIndex + 1; i < this.items.length; ++i) {
            const childEl = this.items[i] as HTMLElement;

            if (!this.isInViewport(containerRect, childEl.getBoundingClientRect())) {
                return i - 1;
            }
        }
        return this.items.length - 1;
    }

    private isInViewport(container: DOMRect, rect: DOMRect) {
        return (
            rect.top >= container.top &&
            rect.left >= container.left &&
            rect.bottom <= container.bottom &&
            rect.right <= container.right
        );
    }

    private updateArrows(): void {
        const firstVisible = this.firstVisibleChild();
        const lastVisible = firstVisible !== null ? this.lastVisibleChildFrom(firstVisible) : null;
        this.updateArrowsVisibility(!!firstVisible, lastVisible !== null && lastVisible !== this.items.length - 1);

    }

    private updateArrowsVisibility(isFirstVisible: boolean, isLastVisible: boolean): void {
        if (!isFirstVisible) {
            this.leftArrowEl.classList.add('hidden');
        } else {
            this.leftArrowEl.classList.remove('hidden');
        }
        if (!isLastVisible) {
            this.rightArrowEl.classList.add('hidden');
        } else {
            this.rightArrowEl.classList.remove('hidden');
        }
    }
}

if (customElements.get("app-netflix-horizontal-list") === undefined) {
    customElements.define("app-netflix-horizontal-list", NetflixHorizontalListComponent);
}
