# imagepool
Simple server and client on GO+gRPC to file exchange

client push <C:client/> <pic.png>       recieve file on dir to server
client get <C:client/> <pic.png>        save file from server on dir
client list                             list of files (on Unix may or may not throw an exception, see server/main.go(155) ) 

There is no to stop server command, but can do,  its no hard
Server create storage folder on .exe file dir 
