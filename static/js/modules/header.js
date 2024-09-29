document.addEventListener("alpine:init", () => {
  Alpine.data("header", () => ({
    show: true,
    lastScrollTop: 0,
    isInteracting: false,
    scrollThreshold: 0,
    init() {
      this.updateThreshold();
      this.$el.addEventListener("mouseenter", () => this.isInteracting = true);
      this.$el.addEventListener("mouseleave", () => this.isInteracting = false);
      window.addEventListener("scroll", this.handleScroll.bind(this), { passive: true });
      window.addEventListener("resize", this.updateThreshold.bind(this));
    },
    updateThreshold() {
      const rootFontSize = parseFloat(getComputedStyle(document.documentElement).fontSize);
      this.scrollThreshold = 1.5 * rootFontSize;
    },
    handleScroll() {
      if (this.isInteracting) return;

      const st = window.pageYOffset || document.documentElement.scrollTop;
      const documentHeight = document.documentElement.scrollHeight;
      const windowHeight = window.innerHeight;

      // At the very top of the page
      if (st <= this.scrollThreshold) {
        this.show = true;
      }
      // At the bottom of the page
      else if ((windowHeight + st) >= documentHeight - this.scrollThreshold) {
        this.show = true;
      }
      // Scrolling down
      else if (st > this.lastScrollTop + this.scrollThreshold) {
        this.show = false;
      }
      // Scrolling up
      else if (st < this.lastScrollTop - this.scrollThreshold) {
        this.show = true;
      }

      this.lastScrollTop = st;
    },
  }));
});
