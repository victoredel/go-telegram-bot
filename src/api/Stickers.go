package api

type MaskPosition struct {
	// This object describes the position on faces where a mask should be placed by default.
	point   string  //The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	x_shift float64 //Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	y_shift float64 //Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	scale   float64 //Mask scaling coefficient. For example, 2.0 means double size.
}

type InputSticker struct {
	// This object describes a sticker to be added to a sticker set.
	sticker       string       //The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one using multipart/form-data, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
	emoji_list    []string     //List of 1-20 emoji associated with the sticker
	mask_position MaskPosition //Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	keywords      []string     //Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
}

type Sticker struct {
	// This object represents a sticker.
	file_id           string       //Identifier for this file, which can be used to download or reuse the file
	file_unique_id    string       //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type              string       //Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	width             int          //Sticker width
	height            int          //Sticker height
	is_animated       bool         //True, if the sticker is animated
	is_video          bool         //True, if the sticker is a video sticker
	thumbnail         PhotoSize    //Optional. Sticker thumbnail in the .WEBP or .JPG format
	emoji             string       //Optional. Emoji associated with the sticker
	set_name          string       //Optional. Name of the sticker set to which the sticker belongs
	premium_animation File         //Optional. For premium regular stickers, premium animation for the sticker
	mask_position     MaskPosition //Optional. For mask stickers, the position where the mask should be placed
	custom_emoji_id   string       //Optional. For custom emoji stickers, unique identifier of the custom emoji
	needs_repainting  bool         //Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	file_size         int          //Optional. File size in bytes
}
