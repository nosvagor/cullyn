document.addEventListener("alpine:init", () => {
    Alpine.data("telescope", () => ({
        openModal() {
            const modal = document.getElementById("telescope");
            modal.showModal();
        },
        closeModal() {
            const modal = document.getElementById("telescope");
            modal.close();
        },
        setupKeyboardShortcut() {
            document.addEventListener("keydown", (e) => {
                if (e.ctrlKey && e.key === "k") {
                    e.preventDefault();
                    this.openModal();
                } else if (e.key === "Escape") {
                    this.closeModal();
                }
            });
        },
        handleSearch(event) {
            console.log("Search query:", event.target.value);
        },
        init() {
            this.setupKeyboardShortcut();
        },
    }));
});
