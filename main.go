package main

/*
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⢀⣠⣴⠾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡴⠂⢀⣤⣶⡿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡾⠋⣠⡴⣫⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⡿⣡⡾⢋⡾⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣴⣶⣶⣶⣶⡶⠶⠶
⠐⠲⠶⣶⣶⠶⢶⣷⣶⣾⣟⣛⠛⠛⠛⠶⠶⠶⠶⠦⣤⣀⣀⠀⠀⠀⠀⢀⣾⢟⣿⡿⠋⢀⡾⠃⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⡴⠶⠞⠛⣩⡿⠟⠋⠉⣰⡶⠋⠁⠀⠀⠀
⠀⠀⠀⠀⠉⠛⢦⣄⠀⠀⠀⠉⠙⠛⠳⠶⢤⣄⡀⠀⠀⠈⠙⠛⠷⠦⣤⣾⣷⣾⠿⠶⠺⢿⣷⣤⣶⣤⣤⣤⣤⠴⠶⠟⠛⠉⠀⢀⣠⠶⠛⠁⠀⠀⢠⡾⠋⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠻⣦⡀⠀⠀⠀⠀⠀⠀⠈⠙⠷⣦⣀⢀⣚⣿⣿⣿⣿⣿⣿⠀⠀⠀⠚⣿⣿⣿⣿⣿⣿⣟⣷⣦⡀⠀⢀⣴⠟⠁⠀⠀⠀⠀⢠⡿⠁⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⡄⠀⠰⣄⠀⠀⠀⠀⠀⠈⣿⢟⣩⣿⣿⣿⡿⠟⠋⠀⠀⠀⠀⠉⠙⠛⢿⣿⣿⣯⡉⠈⢻⣶⠟⠁⠀⠀⣰⠀⠀⠀⣾⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⡄⠀⠘⠷⣄⡀⠀⢠⡾⠁⣴⣿⡿⠋⣁⣭⣶⠧⠀⠀⠀⠀⠀⣀⠀⠀⠉⠻⣿⣷⠀⠀⢻⣦⠀⠀⢀⡟⠀⠀⣼⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣄⠀⠀⠀⠉⢛⡿⠀⢸⡿⠏⣠⡾⠛⢻⣦⠀⠀⠀⠀⠀⠀⣽⡟⠷⣦⡀⢨⣿⡆⠀⠀⢻⣄⣠⠟⠁⣀⡾⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢘⣳⣤⣄⣀⣼⠧⠀⣈⠃⣼⠏⠀⠀⠀⢻⡇⠀⠀⠀⠀⣰⠏⠀⠀⠈⢿⣆⠈⠁⠀⢀⣘⣏⣀⣤⣾⣋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⢾⣿⠿⠅⠀⠀⠁⠀⠀⠉⠛⠿⣶⠲⣶⣶⣾⣷⣤⠀⠀⣤⣿⣤⣴⣶⡶⠶⣿⣤⠶⠛⠋⠉⠉⠉⠉⠛⣿⣿⣗⠶⠤⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢙⣶⢿⣥⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠹⣯⡙⠋⢹⡄⠀⠀⠈⣿⠉⠉⠉⣠⡶⠋⠁⠀⠀⣀⣤⣴⠶⠚⠛⠉⠁⠉⢳⣦⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠟⠁⠀⣀⣈⣹⣿⡶⠶⣤⣄⣀⠀⠀⠀⠀⠈⠻⣄⣸⣡⣶⣶⣶⣿⡆⢀⡾⠋⣀⣤⠶⠾⣋⣉⣤⣶⠶⠾⢓⣿⣏⣝⣯⣉⣿⣄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠤⣼⡷⣤⣶⠿⣹⠏⢀⣠⠿⣤⣤⣈⣉⡛⠻⠶⢤⣄⣀⣹⠫⣧⣦⣤⡤⠤⢹⣿⣶⣻⣭⠶⠟⠛⠋⠁⢀⣠⡴⠞⠛⢧⠀⠹⡀⢹⡏⢻⣆⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢰⣿⡵⢻⠏⢰⠃⠀⡼⠁⠀⠀⣿⠛⠯⠭⣟⣛⣲⣶⢿⡟⠂⠀⢹⠏⠀⠀⠀⠻⣯⣁⣠⡤⠤⠤⠖⠛⠋⠹⢧⡀⠀⠈⢧⠀⡇⠈⡟⢻⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⠋⠀⣿⠀⠘⢦⣴⡃⠀⠀⠐⠙⡆⠀⠀⠀⠀⠀⠀⣼⠀⠀⠀⠘⡆⠀⠀⠀⢠⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⢀⣤⠞⠀⣇⣸⠇⠈⠁⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠶⠺⣿⡍⠳⠦⣤⠴⠶⢺⠉⠉⠀⠀⠀⠀⢹⡀⠀⠀⠀⢻⠀⠀⠀⡼⠀⠀⠒⠦⠤⠤⠤⣴⠶⠞⠋⠻⠓⢲⡶⠛⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡀⠀⠀⠀⠀⢸⣀⣴⡄⠀⠀⠀⠀⢷⡀⠰⢍⠉⠙⠂⢠⣇⡴⢶⠀⠀⠀⠀⢠⡇⠀⠀⠀⠀⠀⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⠀⠀⠀⠀⢸⡿⠀⢻⣴⠟⢧⣤⣼⣧⡀⠀⠙⠀⠀⣸⡿⠠⢸⣧⡼⣆⣠⣾⠁⠀⠀⠀⠀⢸⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⡆⠀⠀⠀⣺⡇⠀⠙⣿⡄⡘⣿⠀⢹⣿⠓⠒⠒⢚⣿⡇⠀⠘⣿⠟⣿⠟⣿⠀⢰⠀⠀⠀⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⠀⠀⢀⡿⠀⠀⠀⠛⠒⠃⠉⠀⠀⢻⡗⠒⠚⠋⣹⡇⠘⠀⠀⠀⠉⢀⣿⡄⠸⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⡆⠀⣼⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢷⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠈⠸⣿⣶⡀⠀⠀⢹⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⡇⣸⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣧⠀⣸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣦⡀⠘⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣷⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣰⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⡟⣷⡀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣇⠘⠿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠶⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠋⢠⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⠦⣄⣀⣀⣀⣀⣀⣀⣀⣀⣀⣠⠤⠾⠛⠶⠤⣄⣀⣀⣀⣀⣀⣀⣀⣄⣤⠤⠴⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*/

import (
	"main.go/app"
)

func main() {
	//db := database.StartDatabase(false)
	app.StartServer()
}
