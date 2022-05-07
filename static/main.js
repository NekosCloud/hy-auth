function logout() {
    fetch("/auth/logout", { method: "DELETE" });
}