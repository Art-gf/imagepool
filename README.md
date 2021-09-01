# imagepool
Simple server and client on GO+gRPC to file exchange

Recieve file from dir to server:
>client push <C:/client/> <pic.png>

Save file from server to dir:
>client get <C:/client/> <pic.png>       

List of files (on Unix may or may not throw an exception, see server/main.go(155)):
>client list

There is no to stop server command, but can do,  its no hard.

Server will create or open storage folder on .exe dir 
