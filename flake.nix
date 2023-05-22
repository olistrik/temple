# https://github.com/numtide/devshell
{
  description = "nix is love, nix is life";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, flake-utils, nixpkgs }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in rec {
        packages = {
          temple = pkgs.buildGoModule {
            name = "temple";
            src = ./.;
            vendorSha256 = null;
          };
        };

        apps = {
          temple = flake-utils.lib.mkApp {
            drv = packages.temple;
            exePath = /bin/temple;
          };
        };

        defaultPackage = packages.temple;
        defaultApp = apps.temple;
      });
}
