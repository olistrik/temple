# https://github.com/numtide/devshell
{
  description = "nix is love, nix is life";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, flake-utils, nixpkgs }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in rec {
        packages = {
          default = pkgs.buildGoModule {
            name = "temple";
            src = ./.;
            vendorSha256 = null;
          };
        };

        apps = {
          default = flake-utils.lib.mkApp {
            drv = packages.temple;
            exePath = /bin/temple;
          };
        };
      }) // {
        overlays.default = (final: prev: rec {
          temple = self.packages."${final.system}".default;
        });
      };
}
