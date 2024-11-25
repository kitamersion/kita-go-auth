with (import <nixpkgs> {config.allowUnfree = true;});
let
in mkShell {

  shellHook = ''
    export PATH="$PATH:$(go env GOPATH)/bin"
  	echo "Done!"
    '';

  buildInputs = [
      go
      postman
  ];

}
