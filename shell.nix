with (import <nixpkgs> {config.allowUnfree = true;});
let
in mkShell {

  shellHook = ''
  	echo "Done!"
    '';

  buildInputs = [
      go
      postman
  ];

}
