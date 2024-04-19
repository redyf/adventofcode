{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    cargo # Downloads your Rust project's dependencies and builds your project
    clippy # A bunch of lints to catch common mistakes and improve your Rust code
    rust-analyzer # Lsp for rust
    rustc # Compiler for rust
    rustfmt # Formatter for rust language
    # rustup # The Rust toolchain installer
  ];
  shellHook = ''
    echo "Environment is ready" | ${pkgs.lolcat}/bin/lolcat;
  '';
}
