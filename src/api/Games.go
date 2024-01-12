package api

type Game struct {
	// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
	title         string          //Title of the game
	description   string          //Description of the game
	photo         []PhotoSize     //Photo that will be displayed in the game message in chats.
	text          string          //Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	text_entities []MessageEntity //Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	animation     Animation
}

type GameHighScore struct {
	// This object represents one row of the high scores table for a game.
	position int  //Position in high score table for the game
	user     User //User
	score    int  //Score
}
