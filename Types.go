package api

type User struct {
	// This object represents a Telegram user or bot.
	// Field 	Type 	Description
	id                          int64   //Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a 64-bit int64eger or double-precision float type are safe for storing this identifier.
	is_bot                      bool    //True, if this user is a bot
	first_name                  string  //User's or bot's first name
	last_name                   *string //Optional. User's or bot's last name
	username                    *string //Optional. User's or bot's username
	language_code               *string //Optional. IETF language tag of the user's language
	is_premium                  *bool   //Optional. True, if this user is a Telegram Premium user
	added_to_attachment_menu    *bool   //Optional. True, if this user added the bot to the attachment menu
	can_join_groups             *bool   //Optional. True, if the bot can be invited to groups. Returned only in getMe.
	can_read_all_group_messages *bool   //Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	supports_inline_queries     *bool   //Optional. True, if the bot supports inline queries. Returned only in getMe.
}

type Chat struct {
	// This object represents a chat.
	// Field 	Type 	Description
	id                                      int64            //Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this identifier.
	Type                                    string           //Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	title                                   *string          //Optional. Title, for supergroups, channels and group chats
	username                                *string          //Optional. Username, for private chats, supergroups and channels if available
	first_name                              *string          //Optional. First name of the other party in a private chat
	last_name                               *string          //Optional. Last name of the other party in a private chat
	is_forum                                *bool            //Optional. True, if the supergroup chat is a forum (has topics enabled)
	photo                                   *ChatPhoto       //Optional. Chat photo. Returned only in getChat.
	active_usernames                        *[]string        //Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels. Returned only in getChat.
	available_reactions                     *[]ReactionType  //Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed. Returned only in getChat.
	accent_color_id                         *int64           //Optional. Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See accent colors for more details. Returned only in getChat. Always returned in getChat.
	background_custom_emoji_id              *string          //Optional. Custom emoji identifier of emoji chosen by the chat for the reply header and link preview background. Returned only in getChat.
	profile_accent_color_id                 *int64           //Optional. Identifier of the accent color for the chat's profile background. See profile accent colors for more details. Returned only in getChat.
	profile_background_custom_emoji_id      *string          //Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background. Returned only in getChat.
	emoji_status_custom_emoji_id            *string          //Optional. Custom emoji identifier of the emoji status of the chat or the other party in a private chat. Returned only in getChat.
	emoji_status_expiration_date            *int64           //Optional. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any. Returned only in getChat.
	bio                                     *string          //Optional. Bio of the other party in a private chat. Returned only in getChat.
	has_private_forwards                    *bool            //Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	has_restricted_voice_and_video_messages *bool            //Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat. Returned only in getChat.
	join_to_send_messages                   *bool            //Optional. True, if users need to join the supergroup before they can send messages. Returned only in getChat.
	join_by_request                         *bool            //Optional. True, if all users directly joining the supergroup need to be approved by supergroup administrators. Returned only in getChat.
	description                             *string          //Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	invite_link                             *string          //Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	pinned_message                          *Message         //Optional. The most recent pinned message (by sending date). Returned only in getChat.
	permissions                             *ChatPermissions //Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	slow_mode_delay                         *int64           //Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	message_auto_delete_time                *int64           //Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	has_aggressive_anti_spam_enabled        *bool            //Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators. Returned only in getChat.
	has_hidden_members                      *bool            //Optional. True, if non-administrators can only get the list of bots and administrators in the chat. Returned only in getChat.
	has_protected_content                   *bool            //Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	has_visible_history                     *bool            //Optional. True, if new chat members will have access to old messages; available only to chat administrators. Returned only in getChat.
	sticker_set_name                        *string          //Optional. For supergroups, name of group sticker set. Returned only in getChat.
	can_set_sticker_set                     *bool            //Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	linked_chat_id                          *int64           //Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it is smaller than 52 bits, so a signed 64 bit int64eger or double-precision float type are safe for storing this identifier. Returned only in getChat.
	location                                *ChatLocation    //Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

type Message struct {
	// This object represents a message.
	// Field 	Type 	Description
	message_id                        int64                          //Unique message identifier inside this chat
	message_thread_id                 *int64                         //Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	from                              *User                          //Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	sender_chat                       *Chat                          //Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field from contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	date                              int64                          //Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	chat                              Chat                           //Chat the message belongs to
	forward_origin                    *MessageOrigin                 //Optional. Information about the original message for forwarded messages
	is_topic_message                  *bool                          //Optional. True, if the message is sent to a forum topic
	is_automatic_forward              *bool                          //Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	reply_to_message                  *Message                       //Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	external_reply                    *ExternalReplyInfo             //Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	quote                             *TextQuote                     //Optional. For replies that quote part of the original message, the quoted part of the message
	via_bot                           *User                          //Optional. Bot through which the message was sent
	edit_date                         *int64                         //Optional. Date the message was last edited in Unix time
	has_protected_content             *bool                          //Optional. True, if the message can't be forwarded
	media_group_id                    *string                        //Optional. The unique identifier of a media message group this message belongs to
	author_signature                  *string                        //Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	text                              *string                        //Optional. For text messages, the actual UTF-8 text of the message
	entities                          *[]MessageEntity               //Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	link_preview_options              *LinkPreviewOptions            //Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	animation                         *Animation                     //Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	audio                             *Audio                         //Optional. Message is an audio file, information about the file
	document                          *Document                      //Optional. Message is a general file, information about the file
	photo                             *[]PhotoSize                   //Optional. Message is a photo, available sizes of the photo
	sticker                           *Sticker                       //Optional. Message is a sticker, information about the sticker
	video                             *Video                         //Optional. Message is a video, information about the video
	video_note                        *VideoNote                     //Optional. Message is a video note, information about the video message
	voice                             *Voice                         //Optional. Message is a voice message, information about the file
	caption                           *string                        //Optional. Caption for the animation, audio, document, photo, video or voice
	caption_entities                  *[]MessageEntity               //Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	has_media_spoiler                 *bool                          //Optional. True, if the message media is covered by a spoiler animation
	contact                           *Contact                       //Optional. Message is a shared contact, information about the contact
	dice                              *Dice                          //Optional. Message is a dice with random value
	game                              *Game                          //Optional. Message is a game, information about the game. More about games »
	poll                              *Poll                          //Optional. Message is a native poll, information about the poll
	venue                             *Venue                         //Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	location                          *Location                      //Optional. Message is a shared location, information about the location
	new_chat_members                  *[]User                        //Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	left_chat_member                  *User                          //Optional. A member was removed from the group, information about them (this member may be the bot itself)
	new_chat_title                    *string                        //Optional. A chat title was changed to this value
	new_chat_photo                    *[]PhotoSize                   //Optional. A chat photo was change to this value
	delete_chat_photo                 *bool                          //Optional. Service message: the chat photo was deleted
	group_chat_created                *bool                          //Optional. Service message: the group has been created
	supergroup_chat_created           *bool                          //Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	channel_chat_created              *bool                          //Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	message_auto_delete_timer_changed *MessageAutoDeleteTimerChanged //Optional. Service message: auto-delete timer settings changed in the chat
	migrate_to_chat_id                *int64                         //Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this identifier.
	migrate_from_chat_id              *int64                         //Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this identifier.
	pinned_message                    *MaybeInaccessibleMessage      //Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	invoice                           *Invoice                       //Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	successful_payment                *SuccessfulPayment             //Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	users_shared                      *UsersShared                   //Optional. Service message: users were shared with the bot
	chat_shared                       *ChatShared                    //Optional. Service message: a chat was shared with the bot
	connected_website                 *string                        //Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	write_access_allowed              *WriteAccessAllowed            //Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	passport_data                     *PassportData                  //Optional. Telegram Passport data
	proximity_alert_triggered         *ProximityAlertTriggered       //Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	forum_topic_created               *ForumTopicCreated             //Optional. Service message: forum topic created
	forum_topic_edited                *ForumTopicEdited              //Optional. Service message: forum topic edited
	giveaway                          *Giveaway                      //Optional. The message is a scheduled giveaway message
	giveaway_winners                  *GiveawayWinners               //Optional. A giveaway with public winners was completed
	giveaway_completed                *GiveawayCompleted             //Optional. Service message: a giveaway without public winners was completed
	video_chat_scheduled              *VideoChatScheduled            //Optional. Service message: video chat scheduled
	video_chat_ended                  *VideoChatEnded                //Optional. Service message: video chat ended
	video_chat_participants_invited   *VideoChatParticipantsInvited  //Optional. Service message: new participants invited to a video chat
	web_app_data                      *WebAppData                    //Optional. Service message: data sent by a Web App
	reply_markup                      *InlineKeyboardMarkup          //Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	// forum_topic_closed                *ForumTopicClosed              //Optional. Service message: forum topic closed
	// forum_topic_reopened              *ForumTopicReopened            //Optional. Service message: forum topic reopened
	// general_forum_topic_hidden        *GeneralForumTopicHidden       //Optional. Service message: the 'General' forum topic hidden
	// general_forum_topic_unhidden      *GeneralForumTopicUnhidden     //Optional. Service message: the 'General' forum topic unhidden
	// giveaway_created                  *GiveawayCreated               //Optional. Service message: a scheduled giveaway was created
	// video_chat_started                *VideoChatStarted              //Optional. Service message: video chat started
}

type MessageId struct {
	// This object represents a unique message identifier.
	// Field 	Type 	Description
	message_id int64 //Unique message identifier
}

type InaccessibleMessage struct {
	// This object describes a message that was deleted or is otherwise inaccessible to the bot.
	// Field 	Type 	Description
	chat       Chat  //Chat the message belonged to
	message_id int64 //Unique message identifier inside the chat
	date       int64 //Always 0. The field can be used to differentiate regular and inaccessible messages.
}

type MaybeInaccessibleMessage struct {
	// This object describes a message that can be inaccessible to the bot. It can be one of
	message             Message
	inaccessibleMessage InaccessibleMessage
}

type MessageEntity struct {
	// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
	// Field 	Type 	Description
	Type            string  //Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	offset          int64   //Offset in UTF-16 code units to the start of the entity
	length          int64   //Length of the entity in UTF-16 code units
	url             *string //Optional. For “text_link” only, URL that will be opened after user taps on the text
	user            *User   //Optional. For “text_mention” only, the mentioned user
	language        *string //Optional. For “pre” only, the programming language of the entity text
	custom_emoji_id *string //Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

type TextQuote struct {
	// This object contains information about the quoted part of a message that is replied to by the given message.
	// Field 	Type 	Description
	text      string           //Text of the quoted part of a message that is replied to by the given message
	entities  *[]MessageEntity //Optional. Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	position  int64            //Approximate quote position in the original message in UTF-16 code units as specified by the sender
	is_manual *bool            //Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
}

type ExternalReplyInfo struct {

	// This object contains information about a message that is being replied to, which may come from another chat or forum topic.
	// Field 	Type 	Description
	origin               MessageOrigin       //Origin of the message replied to by the given message
	chat                 *Chat               //Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	message_id           *int64              //Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	link_preview_options *LinkPreviewOptions //Optional. Options used for link preview generation for the original message, if it is a text message
	animation            *Animation          //Optional. Message is an animation, information about the animation
	audio                *Audio              //Optional. Message is an audio file, information about the file
	document             *Document           //Optional. Message is a general file, information about the file
	photo                *[]PhotoSize        //Optional. Message is a photo, available sizes of the photo
	sticker              *Sticker            //Optional. Message is a sticker, information about the sticker
	// story                *Story              //Optional. Message is a forwarded story
	video             *Video           //Optional. Message is a video, information about the video
	video_note        *VideoNote       //Optional. Message is a video note, information about the video message
	voice             *Voice           //Optional. Message is a voice message, information about the file
	has_media_spoiler *bool            //Optional. True, if the message media is covered by a spoiler animation
	contact           *Contact         //Optional. Message is a shared contact, information about the contact
	dice              *Dice            //Optional. Message is a dice with random value
	game              *Game            //Optional. Message is a game, information about the game. More about games »
	giveaway          *Giveaway        //Optional. Message is a scheduled giveaway, information about the giveaway
	giveaway_winners  *GiveawayWinners //Optional. A giveaway with public winners was completed
	invoice           *Invoice         //Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	location          *Location        //Optional. Message is a shared location, information about the location
	poll              *Poll            //Optional. Message is a native poll, information about the poll
	venue             *Venue           //Optional. Message is a venue, information about the venue
}
type ReplyParameters struct {
	// Describes reply parameters for the message that is being sent.
	// Field 	Type 	Description
	message_id                  int64            //Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	chat_id                     *int64           //Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername)
	allow_sending_without_reply *bool            //Optional. Pass True if the message should be sent even if the specified message to be replied to is not found; can be used only for replies in the same chat and forum topic.
	quote                       *string          //Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
	quote_parse_mode            *string          //Optional. Mode for parsing entities in the quote. See formatting options for more details.
	quote_entities              *[]MessageEntity //Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
	quote_position              *int64           //Optional. Position of the quote in the original message in UTF-16 code units
}

type MessageOrigin struct {
	// This object describes the origin of a message. It can be one of
	Type             string  //Type of the message origin, “user, hidden_user, chat, channel”
	date             int64   //Date the message was sent originally in Unix time
	sender_user      User    //User that sent the message originally
	sender_user_name string  //Name of the user that sent the message originally
	sender_chat      Chat    //Chat that sent the message originally
	author_signature *string //Optional. For messages originally sent by an anonymous chat administrator, original message author signature
	message_id       int64   //Unique message identifier inside the chat
}

type PhotoSize struct {
	// This object represents one size of a photo or a file / sticker thumbnail.
	// Field 	Type 	Description
	file_id        string //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	width          int64  //Photo width
	height         int64  //Photo height
	file_size      *int64 //Optional. File size in bytes
}

type Animation struct {
	// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
	// Field 	Type 	Description
	file_id        string     //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string     //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	width          int64      //Video width as defined by sender
	height         int64      //Video height as defined by sender
	duration       int64      //Duration of the video in seconds as defined by sender
	thumbnail      *PhotoSize //Optional. Animation thumbnail as defined by sender
	file_name      *string    //Optional. Original animation filename as defined by sender
	mime_type      *string    //Optional. MIME type of the file as defined by sender
	file_size      *int64     //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
}

type Audio struct {
	// This object represents an audio file to be treated as music by the Telegram clients.
	// Field 	Type 	Description
	file_id        string     //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string     //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	duration       int64      //Duration of the audio in seconds as defined by sender
	performer      *string    //Optional. Performer of the audio as defined by sender or by audio tags
	title          *string    //Optional. Title of the audio as defined by sender or by audio tags
	file_name      *string    //Optional. Original filename as defined by sender
	mime_type      *string    //Optional. MIME type of the file as defined by sender
	file_size      *int64     //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
	thumbnail      *PhotoSize //Optional. Thumbnail of the album cover to which the music file belongs
}

type Document struct {
	// This object represents a general file (as opposed to photos, voice messages and audio files).
	// Field 	Type 	Description
	file_id        string     //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string     //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	thumbnail      *PhotoSize //Optional. Document thumbnail as defined by sender
	file_name      *string    //Optional. Original filename as defined by sender
	mime_type      *string    //Optional. MIME type of the file as defined by sender
	file_size      *int64     //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
}

// type Story struct{
// // This object represents a message about a forwarded story in the chat. Currently holds no information.
// }

type Video struct {
	// This object represents a video file.
	// Field 	Type 	Description
	file_id        string     //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string     //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	width          int64      //Video width as defined by sender
	height         int64      //Video height as defined by sender
	duration       int64      //Duration of the video in seconds as defined by sender
	thumbnail      *PhotoSize //Optional. Video thumbnail
	file_name      *string    //Optional. Original filename as defined by sender
	mime_type      *string    //Optional. MIME type of the file as defined by sender
	file_size      *int64     //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
}

type VideoNote struct {
	// This object represents a video message (available in Telegram apps as of v.4.0).
	// Field 	Type 	Description
	file_id        string     //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string     //	Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	length         int64      //Video width and height (diameter of the video message) as defined by sender
	duration       int64      //Duration of the video in seconds as defined by sender
	thumbnail      *PhotoSize //Optional. Video thumbnail
	file_size      *int64     //Optional. File size in bytes
}

type Voice struct {
	// This object represents a voice note.
	// Field 	Type 	Description
	file_id        string  //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string  //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	duration       int64   //Duration of the audio in seconds as defined by sender
	mime_type      *string //Optional. MIME type of the file as defined by sender
	file_size      *int64  //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
}

type Contact struct {
	// This object represents a phone contact.
	// Field 	Type 	Description
	phone_number string  //Contact's phone number
	first_name   string  //Contact's first name
	last_name    *string //Optional. Contact's last name
	user_id      *int64  //Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a 64-bit int64eger or double-precision float type are safe for storing this identifier.
	vcard        *string //Optional. Additional data about the contact in the form of a vCard
}

type Dice struct {
	// This object represents an animated emoji that displays a random value.
	// Field 	Type 	Description
	emoji string //Emoji on which the dice throw animation is based
	value int64  //Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
}

type PollOption struct {
	// This object contains information about one answer option in a poll.
	// Field 	Type 	Description
	text        string //Option text, 1-100 characters
	voter_count int64  //Number of users that voted for this option
}

type PollAnswer struct {
	// This object represents an answer of a user in a non-anonymous poll.
	// Field 	Type 	Description
	poll_id    string  //Unique poll identifier
	voter_chat *Chat   //Optional. The chat that changed the answer to the poll, if the voter is anonymous
	user       *User   //Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	option_ids []int64 //0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}

type Poll struct {
	// This object contains information about a poll.
	// Field 	Type 	Description
	id                      string           //Unique poll identifier
	question                string           //Poll question, 1-300 characters
	options                 []PollOption     //List of poll options
	total_voter_count       int64            //Total number of users that voted in the poll
	is_closed               bool             //True, if the poll is closed
	is_anonymous            bool             //True, if the poll is anonymous
	Type                    string           //Poll type, currently can be “regular” or “quiz”
	allows_multiple_answers bool             //True, if the poll allows multiple answers
	correct_option_id       *int64           //Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	explanation             *string          //Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	explanation_entities    *[]MessageEntity //Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	open_period             *int64           //Optional. Amount of time in seconds the poll will be active after creation
	close_date              *int64           //Optional. Point64 in time (Unix timestamp) when the poll will be automatically closed
}

type Location struct {
	// This object represents a point64 on the map.
	// Field 	Type 	Description
	longitude              float64  //Longitude as defined by sender
	latitude               float64  //Latitude as defined by sender
	horizontal_accuracy    *float64 //Optional. The radius of uncertaint64y for the location, measured in meters; 0-1500
	live_period            *int64   //Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	heading                *int64   //Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	proximity_alert_radius *int64   //Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

type Venue struct {
	// This object represents a venue.
	// Field 	Type 	Description
	location          Location //Venue location. Can't be a live location
	title             string   //Name of the venue
	address           string   //Address of the venue
	foursquare_id     *string  //Optional. Foursquare identifier of the venue
	foursquare_type   *string  //Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	google_place_id   *string  //Optional. Google Places identifier of the venue
	google_place_type *string  //Optional. Google Places type of the venue. (See supported types.)
}

type WebAppData struct {
	// Describes data sent from a Web App to the bot.
	// Field 	Type 	Description
	data        string //The data. Be aware that a bad client can send arbitrary data in this field.
	button_text string //Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
}

type ProximityAlertTriggered struct {
	// This object represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
	// Field 	Type 	Description
	traveler User  //User that triggered the alert
	watcher  User  //User that set the alert
	distance int64 //The distance between the users
}

type MessageAutoDeleteTimerChanged struct {
	// This object represents a service message about a change in auto-delete timer settings.
	// Field 	Type 	Description
	message_auto_delete_time int64 //New auto-delete time for messages in the chat; in seconds
}

type ForumTopicCreated struct {
	// This object represents a service message about a new forum topic created in the chat.
	// Field 	Type 	Description
	name                 string  //Name of the topic
	icon_color           int64   //Color of the topic icon in RGB format
	icon_custom_emoji_id *string //Optional. Unique identifier of the custom emoji shown as the topic icon
}

// type ForumTopicClosed struct{
// // This object represents a service message about a forum topic closed in the chat. Currently holds no information.
// }

type ForumTopicEdited struct {
	// This object represents a service message about an edited forum topic.
	// Field 	Type 	Description
	name                 *string //Optional. New name of the topic, if it was edited
	icon_custom_emoji_id *string //Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
}

// type ForumTopicReopened struct{
// // This object represents a service message about a forum topic reopened in the chat. Currently holds no information.
// }

// type GeneralForumTopicHidden struct{
// // This object represents a service message about General forum topic hidden in the chat. Currently holds no information.
// }

// type GeneralForumTopicUnhidden struct{
// // This object represents a service message about General forum topic unhidden in the chat. Currently holds no information.
// }

type UsersShared struct {
	// This object contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
	// Field 	Type 	Description
	request_id int64   //Identifier of the request
	user_ids   []int64 //Identifiers of the shared users. These numbers may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting them. But they have at most 52 significant bits, so 64-bit int64egers or double-precision float types are safe for storing these identifiers. The bot may not have access to the users and could be unable to use these identifiers, unless the users are already known to the bot by some other means.
}

type ChatShared struct {
	// This object contains information about the chat whose identifier was shared with the bot using a KeyboardButtonRequestChat button.
	// Field 	Type 	Description
	request_id int64 //Identifier of the request
	chat_id    int64 //Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a 64-bit int64eger or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
}

type WriteAccessAllowed struct {
	// This object represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
	// Field 	Type 	Description
	from_request         *bool   //Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	web_app_name         *string //Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	from_attachment_menu *bool   //Optional. True, if the access was granted when the bot was added to the attachment or side menu
}

type VideoChatScheduled struct {
	// This object represents a service message about a video chat scheduled in the chat.
	// Field 	Type 	Description
	start_date int64 //Point64 in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

// type VideoChatStarted struct{
// // This object represents a service message about a video chat started in the chat. Currently holds no information.
// }

type VideoChatEnded struct {
	// This object represents a service message about a video chat ended in the chat.
	// Field 	Type 	Description
	duration int64 //Video chat duration in seconds
}

type VideoChatParticipantsInvited struct {
	// This object represents a service message about new members invited to a video chat.
	// Field 	Type 	Description
	users []User //New members that were invited to the video chat
}

// type GiveawayCreated struct{
// // This object represents a service message about the creation of a scheduled giveaway. Currently holds no information.
// }

type Giveaway struct {
	// This object represents a message about a scheduled giveaway.
	// Field 	Type 	Description
	chats                            []Chat    //The list of chats which the user must join to participate in the giveaway
	winners_selection_date           int64     //Point64 in time (Unix timestamp) when winners of the giveaway will be selected
	winner_count                     int64     //The number of users which are supposed to be selected as winners of the giveaway
	only_new_members                 *bool     //Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	has_public_winners               *bool     //Optional. True, if the list of giveaway winners will be visible to everyone
	prize_description                *string   //Optional. Description of additional giveaway prize
	country_codes                    *[]string //Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	premium_subscription_month_count *int64    //Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
}

type GiveawayWinners struct {
	// This object represents a message about the completion of a giveaway with public winners.
	// Field 	Type 	Description
	chat                             Chat    //The chat that created the giveaway
	giveaway_message_id              int64   //Identifier of the messsage with the giveaway in the chat
	winners_selection_date           int64   //Point64 in time (Unix timestamp) when winners of the giveaway were selected
	winner_count                     int64   //Total number of winners in the giveaway
	winners                          []User  //List of up to 100 winners of the giveaway
	additional_chat_count            *int64  //Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	premium_subscription_month_count *int64  //Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
	unclaimed_prize_count            *int64  //Optional. Number of undistributed prizes
	only_new_members                 *bool   //Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	was_refunded                     *bool   //Optional. True, if the giveaway was canceled because the payment for it was refunded
	prize_description                *string //Optional. Description of additional giveaway prize
}

type GiveawayCompleted struct {
	// This object represents a service message about the completion of a giveaway without public winners.
	// Field 	Type 	Description
	winner_count          int64    //Number of winners in the giveaway
	unclaimed_prize_count *int64   //Optional. Number of undistributed prizes
	giveaway_message      *Message //Optional. Message with the giveaway that was completed, if it wasn't deleted
}

type LinkPreviewOptions struct {
	// Describes the options used for link preview generation.
	// Field 	Type 	Description
	is_disabled        *bool   //Optional. True, if the link preview is disabled
	url                *string //Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	prefer_small_media *bool   //Optional. True, if the media in the link preview is suppposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	prefer_large_media *bool   //Optional. True, if the media in the link preview is suppposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	show_above_text    *bool   //Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
}

type UserProfilePhotos struct {
	// This object represent a user's profile pictures.
	// Field 	Type 	Description
	total_count int64         //Total number of profile pictures the target user has
	photos      [][]PhotoSize //Requested profile pictures (in up to 4 sizes each)
}

type File struct {
	// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
	// The maximum file size to download is 20 MB
	// Field 	Type 	Description
	file_id        string  //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string  //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	file_size      *int64  //Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this value.
	file_path      *string //Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

type WebAppInfo struct {
	// Describes a Web App.
	// Field 	Type 	Description
	url string //An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

type ReplyKeyboardMarkup struct {
	// This object represents a custom keyboard with reply options (see int64roduction to bots for details and examples).
	// Field 	Type 	Description
	keyboard                [][]KeyboardButton //[]button rows, each represented by an []KeyboardButton objects
	is_persistent           *bool              //Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	resize_keyboard         *bool              //Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	one_time_keyboard       *bool              //Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	input_field_placeholder *string            //Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	selective               *bool              //Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message. Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

type KeyboardButton struct {
	// This object represents one button of the reply keyboard. For simple text buttons, string can be used instead of this object to specify the button *text. The optional fields web_app, request_users, request_chat, request_contact, request_location, and request_poll are mutually exclusive.
	// Field 	Type 	Description
	text             string                      //Text of the button. If none *of the optional fields are used, it will be sent as a message when the button is pressed
	request_users    *KeyboardButtonRequestUsers //Optional. If specified, pressing the button will open a list of suitable users. Identifiers of selected users will be sent to the bot in a “users_shared” service message. Available in private chats only.
	request_chat     *KeyboardButtonRequestChat  //Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	request_contact  *bool                       //Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	request_location *bool                       //Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	request_poll     *KeyboardButtonPollType     //Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	web_app          *WebAppInfo                 //Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

//Note: request_users and request_chat options will only work in Telegram versions released after 3 February, 2023. Older clients will display unsupported message.

type KeyboardButtonRequestUsers struct {
	// This object defines the criteria used to request suitable users. The identifiers of the selected users will be shared with the bot when the corresponding button is pressed. More about requesting users »
	// Field 	Type 	Description
	request_id      int64  //Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	user_is_bot     *bool  //Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	user_is_premium *bool  //Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	max_quantity    *int64 //Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
}

type KeyboardButtonRequestChat struct {
	// This object defines the criteria used to request a suitable chat. The identifier of the selected chat will be shared with the bot when the corresponding button is pressed. More about requesting chats »
	// Field 	Type 	Description
	request_id                int64                    //Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	chat_is_channel           bool                     //Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	chat_is_forum             *bool                    //Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	chat_has_username         *bool                    //Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	chat_is_created           *bool                    //Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	user_administrator_rights *ChatAdministratorRights //Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	bot_administrator_rights  *ChatAdministratorRights //Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	bot_is_member             *bool                    //Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
}

type KeyboardButtonPollType struct {
	// This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
	// Field 	Type 	Description
	Type *string //Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

type ReplyKeyboardRemove struct {
	// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
	// Field 	Type 	Description
	remove_keyboard bool  //Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	selective       *bool //Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message. Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

type InlineKeyboardMarkup struct {
	// This object represents an inline keyboard that appears right next to the message it belongs to.
	// Field 	Type 	Description
	inline_keyboard [][]InlineKeyboardButton //[]button rows, each represented by an []InlineKeyboardButton objects
}

type InlineKeyboardButton struct {
	// This object represents one button of an inline keyboard. You must use exactly one *of the optional fields.
	// Field 	Type 	Description
	text                             string                       //Label text on the button
	url                              *string                      //Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	callback_data                    *string                      //Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	web_app                          *WebAppInfo                  //Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
	login_url                        *LoginUrl                    //Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	switch_inline_query              *string                      //Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted.
	switch_inline_query_current_chat *string                      //Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options.
	switch_inline_query_chosen_chat  *SwitchInlineQueryChosenChat //Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field
	// callback_game                    *CallbackGame                //Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
	pay *bool //Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

type LoginUrl struct {
	// This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
	// TITLE
	// Telegram apps support these buttons as of version 5.7.
	//
	//	Sample bot: @discussbot
	//
	// Field 	Type 	Description
	url                  string  //An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data. NOTE: You must always check the hash of the received data to verify the authentication and the int64egrity of the data as described in Checking authorization.
	forward_text         *string //Optional. New text of the button in forwarded messages.
	bot_username         *string //Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	request_write_access *bool   //Optional. Pass True to request the permission for your bot to send messages to the user.
}

type SwitchInlineQueryChosenChat struct {
	// This object represents an inline button that switches the current user to inline mode in a chosen chat, *with an optional default inline query.
	// Field 	Type 	Description
	query               *string //Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	allow_user_chats    *bool   //Optional. True, if private chats with users can be chosen
	allow_bot_chats     *bool   //Optional. True, if private chats with bots can be chosen
	allow_group_chats   *bool   //Optional. True, if group and supergroup chats can be chosen
	allow_channel_chats *bool   //Optional. True, if channel chats can be chosen
}

type CallbackQuery struct {
	// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
	// Field 	Type 	Description
	id                string                    //Unique identifier for this query
	from              User                      //Sender
	message           *MaybeInaccessibleMessage //Optional. Message sent by the bot with the callback button that originated the query
	inline_message_id *string                   //Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	chat_instance     string                    //Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	data              *string                   //Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	game_short_name   *string                   //Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	/*
	   NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any *of the optional parameters).
	*/
}

type ForceReply struct {
	// Upon receiving a message with this object, Telegram clients will display a reply int64erface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step int64erfaces without having to sacrifice privacy mode.
	// Field 	Type 	Description
	force_reply             bool    //Shows reply int64erface to the user, as if they manually selected the bot's message and tapped 'Reply'
	input_field_placeholder *string //Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	selective               *bool   //Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
	/*
	   Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions). There could be two ways to create a new poll:

	   	Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing for hardcore users but lacks modern day polish.
	   	Guide the user through a step-by-step process. 'Please send me your question', 'Cool, now let's add the first answer option', 'Great. Keep adding answer options, then send /done when you're ready'.

	   The last option is definitely more attractive. And if you use ForceReply in your bot's questions, it will receive the user's answers even if it only receives replies, commands and mentions - without any extra work for the user.
	*/
}

type ChatPhoto struct {
	// This object represents a chat photo.
	// Field 	Type 	Description
	small_file_id        string //File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	small_file_unique_id string //Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	big_file_id          string //File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	big_file_unique_id   string //Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

type ChatInviteLink struct {
	// Represents an invite link for a chat.
	// Field 	Type 	Description
	invite_link                string  //The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	creator                    User    //Creator of the link
	creates_join_request       bool    //True, if users joining the chat via the link need to be approved by chat administrators
	is_primary                 bool    //True, if the link is primary
	is_revoked                 bool    //True, if the link is revoked
	name                       *string //Optional. Invite link name
	expire_date                *int64  //Optional. Point64 in time (Unix timestamp) when the link will expire or has been expired
	member_limit               *int64  //Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	pending_join_request_count *int64  //Optional. Number of pending join requests created using this link
}

type ChatAdministratorRights struct {
	// Represents the rights of an administrator in a chat.
	// Field 	Type 	Description
	is_anonymous           bool  //True, if the user's presence in the chat is hidden
	can_manage_chat        bool  //True, if the administrator can access the chat event log, boost list in channels, see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	can_delete_messages    bool  //True, if the administrator can delete messages of other users
	can_manage_video_chats bool  //True, if the administrator can manage video chats
	can_restrict_members   bool  //True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	can_promote_members    bool  //True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appoint64ed by the user)
	can_change_info        bool  //True, if the user is allowed to change the chat title, photo and other settings
	can_invite_users       bool  //True, if the user is allowed to invite new users to the chat
	can_post_messages      *bool //Optional. True, if the administrator can post messages in the channel, or access channel statistics; channels only
	can_edit_messages      *bool //Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	can_pin_messages       *bool //Optional. True, if the user is allowed to pin messages; groups and supergroups only
	can_post_stories       *bool //Optional. True, if the administrator can post stories in the channel; channels only
	can_edit_stories       *bool //Optional. True, if the administrator can edit stories posted by other users; channels only
	can_delete_stories     *bool //Optional. True, if the administrator can delete stories posted by other users; channels only
	can_manage_topics      *bool //Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
}

type ChatMemberUpdated struct {
	// This object represents changes in the status of a chat member.
	// Field 	Type 	Description
	chat                        Chat            //Chat the user belongs to
	from                        User            //Performer of the action, which resulted in the change
	date                        int64           //Date the change was done in Unix time
	old_chat_member             ChatMember      //Previous information about the chat member
	new_chat_member             ChatMember      //New information about the chat member
	invite_link                 *ChatInviteLink //Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	via_chat_folder_invite_link *bool           //Optional. True, if the user joined the chat via a chat folder invite link
}

type ChatMember struct {
	// This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
	/*
		// ChatMemberOwner
		// ChatMemberAdministrator
		// ChatMemberMember
		// ChatMemberRestricted
		// ChatMemberLeft
		// ChatMemberBanned
	*/
	status                    string  //The member's status in the chat: [“creator"|"administrator”|“member”|“restricted”|“left”|“kicked”]
	user                      User    //Information about the user
	is_anonymous              bool    //True, if the user's presence in the chat is hidden
	custom_title              *string //Optional. Custom title for this user
	can_be_edited             bool    //True, if the bot is allowed to edit administrator privileges of that user
	can_manage_chat           bool    //True, if the administrator can access the chat event log, boost list in channels, see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	can_delete_messages       bool    //True, if the administrator can delete messages of other users
	can_manage_video_chats    bool    //True, if the administrator can manage video chats
	can_restrict_members      bool    //True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	can_promote_members       bool    //True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appoint64ed by the user)
	can_change_info           bool    //True, if the user is allowed to change the chat title, photo and other settings
	can_invite_users          bool    //True, if the user is allowed to invite new users to the chat
	can_post_messages         *bool   //Optional. True, if the administrator can post messages in the channel, or access channel statistics; channels only
	can_edit_messages         *bool   //Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	can_pin_messages          *bool   //Optional. True, if the user is allowed to pin messages; groups and supergroups only
	can_post_stories          *bool   //Optional. True, if the administrator can post stories in the channel; channels only
	can_edit_stories          *bool   //Optional. True, if the administrator can edit stories posted by other users; channels only
	can_delete_stories        *bool   //Optional. True, if the administrator can delete stories posted by other users; channels only
	can_manage_topics         *bool   //Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
	is_member                 bool    //True, if the user is a member of the chat at the moment of the request
	can_send_messages         bool    //True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	can_send_audios           bool    //True, if the user is allowed to send audios
	can_send_documents        bool    //True, if the user is allowed to send documents
	can_send_photos           bool    //True, if the user is allowed to send photos
	can_send_videos           bool    //True, if the user is allowed to send videos
	can_send_video_notes      bool    //True, if the user is allowed to send video notes
	can_send_voice_notes      bool    //True, if the user is allowed to send voice notes
	can_send_polls            bool    //True, if the user is allowed to send polls
	can_send_other_messages   bool    //True, if the user is allowed to send animations, games, stickers and use inline bots
	can_add_web_page_previews bool    //True, if the user is allowed to add web page previews to their messages
	until_date                int64   //Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
}

type ChatJoinRequest struct {
	// Represents a join request sent to a chat.
	// Field 	Type 	Description
	chat         Chat            //Chat to which the request was sent
	from         User            //User that sent the join request
	user_chat_id int64           //Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a 64-bit int64eger or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	date         int64           //Date the request was sent in Unix time
	bio          *string         //Optional. Bio of the user.
	invite_link  *ChatInviteLink //Optional. Chat invite link that was used by the user to send the join request
}

type ChatPermissions struct {
	// Describes actions that a non-administrator user is allowed to take in a chat.
	// Field 	Type 	Description
	can_send_messages         *bool //Optional. True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	can_send_audios           *bool //Optional. True, if the user is allowed to send audios
	can_send_documents        *bool //Optional. True, if the user is allowed to send documents
	can_send_photos           *bool //Optional. True, if the user is allowed to send photos
	can_send_videos           *bool //Optional. True, if the user is allowed to send videos
	can_send_video_notes      *bool //Optional. True, if the user is allowed to send video notes
	can_send_voice_notes      *bool //Optional. True, if the user is allowed to send voice notes
	can_send_polls            *bool //Optional. True, if the user is allowed to send polls
	can_send_other_messages   *bool //Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
	can_add_web_page_previews *bool //Optional. True, if the user is allowed to add web page previews to their messages
	can_change_info           *bool //Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	can_invite_users          *bool //Optional. True, if the user is allowed to invite new users to the chat
	can_pin_messages          *bool //Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	can_manage_topics         *bool //Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages
}

type ChatLocation struct {
	// Represents a location to which a chat is connected.
	// Field 	Type 	Description
	location Location //The location to which the supergroup is connected. Can't be a live location.
	address  string   //Location address; 1-64 characters, as defined by the chat owner
}

type ReactionType struct {
	// This object describes the type of a reaction. Currently, it can be one of
	// ReactionTypeEmoji
	// ReactionTypeCustomEmoji
	Type            string //Type of the reaction, always “emoji”
	emoji           string //Reaction emoji. Currently, it can be one of "👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱", "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
	custom_emoji_id string //Custom emoji identifier
}

type ReactionCount struct {
	// Represents a reaction added to a message along with the number of times it was added.
	// Field 	Type 	Description
	Type        ReactionType //Type of the reaction
	total_count int64        //Number of times the reaction was added
}

type MessageReactionUpdated struct {
	// This object represents a change of a reaction on a message performed by a user.
	// Field 	Type 	Description
	chat         Chat           //The chat containing the message the user reacted to
	message_id   int64          //Unique identifier of the message inside the chat
	user         *User          //Optional. The user that changed the reaction, if the user isn't anonymous
	actor_chat   *Chat          //Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	date         int64          //Date of the change in Unix time
	old_reaction []ReactionType //Previous list of reaction types that were set by the user
	new_reaction []ReactionType //New list of reaction types that have been set by the user
}

type MessageReactionCountUpdated struct {
	// This object represents reaction changes on a message with anonymous reactions.
	// Field 	Type 	Description
	chat       Chat            //The chat containing the message
	message_id int64           //Unique message identifier inside the chat
	date       int64           //Date of the change in Unix time
	reactions  []ReactionCount //List of reactions that are present on the message
}

type ForumTopic struct {
	// This object represents a forum topic.
	// Field 	Type 	Description
	message_thread_id    int64   //Unique identifier of the forum topic
	name                 string  //Name of the topic
	icon_color           int64   //Color of the topic icon in RGB format
	icon_custom_emoji_id *string //Optional. Unique identifier of the custom emoji shown as the topic icon
}

type BotCommand struct {
	// This object represents a bot command.
	// Field 	Type 	Description
	command     string //Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	description string //Description of the command; 1-256 characters.
}

type BotCommandScope struct {
	Type    string //Scope type, must be ["default"|"all_private_chats"|"all_group_chats"|"all_chat_administrators"|"chat"|"chat_administrators"|"chat_member"]
	chat_id int64  //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	user_id int64  //Unique identifier of the target user
	// This object represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
	//
	//	BotCommandScopeDefault
	//	BotCommandScopeAllPrivateChats
	//	BotCommandScopeAllGroupChats
	//	BotCommandScopeAllChatAdministrators
	//	BotCommandScopeChat
	//	BotCommandScopeChatAdministrators
	//	BotCommandScopeChatMember
	//
	// Determining list of commands
	// The following algorithm is used to determine the list of commands for a particular user viewing the bot menu. The first list of commands which is set is returned:
	// Commands in the chat with the bot
	//
	//	botCommandScopeChat + language_code
	//	botCommandScopeChat
	//	botCommandScopeAllPrivateChats + language_code
	//	botCommandScopeAllPrivateChats
	//	botCommandScopeDefault + language_code
	//	botCommandScopeDefault
	//
	// Commands in group and supergroup chats
	//
	//	botCommandScopeChatMember + language_code
	//	botCommandScopeChatMember
	//	botCommandScopeChatAdministrators + language_code (administrators only)
	//	botCommandScopeChatAdministrators (administrators only)
	//	botCommandScopeChat + language_code
	//	botCommandScopeChat
	//	botCommandScopeAllChatAdministrators + language_code (administrators only)
	//	botCommandScopeAllChatAdministrators (administrators only)
	//	botCommandScopeAllGroupChats + language_code
	//	botCommandScopeAllGroupChats
	//	botCommandScopeDefault + language_code
	//	botCommandScopeDefault
}

type BotName struct {
	// This object represents the bot's name.
	// Field 	Type 	Description
	name string //The bot's name
}

type BotDescription struct {
	// This object represents the bot's description.
	// Field 	Type 	Description
	description string //The bot's description
}

type BotShortDescription struct {
	// This object represents the bot's short description.
	// Field 	Type 	Description
	short_description string //The bot's short description
}

type MenuButton struct {
	// This object describes the bot's menu button in a private chat. It should be one of
	//     MenuButtonCommands
	//     MenuButtonWebApp
	//     MenuButtonDefault
	// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
	Type    string     //Type of the button, must be ["commands"|"web_app"|"default"]
	text    string     //Text on the button
	web_app WebAppInfo //Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
}

type ChatBoostSource struct {
	// This object describes the source of a chat boost. It can be one of
	//     ChatBoostSourcePremium
	//     ChatBoostSourceGiftCode
	//     ChatBoostSourceGiveaway
	// ChatBoostSourcePremium

	// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
	// Field 	Type 	Description
	source              string //Source of the boost, must be: [“premium”|“gift_code”|“giveaway”
	user                User   //User that boosted the chat | User for which the gift code was *created | Optional. User that won the prize in the giveaway if any
	giveaway_message_id int64  //Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	is_unclaimed        *bool  //Optional. True, if the giveaway was completed, but there was no user to win the prize
}

type ChatBoost struct {
	// This object contains information about a chat boost.
	// Field 	Type 	Description
	boost_id        string          //Unique identifier of the boost
	add_date        int64           //Point64 in time (Unix timestamp) when the chat was boosted
	expiration_date int64           //Point64 in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	source          ChatBoostSource //Source of the added boost
}

type ChatBoostUpdated struct {
	// This object represents a boost added to a chat or changed.
	// Field 	Type 	Description
	chat  Chat      //Chat which was boosted
	boost ChatBoost //Infomation about the chat boost
}

type ChatBoostRemoved struct {
	// This object represents a boost removed from a chat.
	// Field 	Type 	Description
	chat        Chat            //Chat which was boosted
	boost_id    string          //Unique identifier of the boost
	remove_date int64           //Point64 in time (Unix timestamp) when the boost was removed
	source      ChatBoostSource //Source of the removed boost
}

type UserChatBoosts struct {
	// This object represents a list of boosts added to a chat by a user.
	// Field 	Type 	Description
	boosts []ChatBoost //The list of boosts added to the chat by the user
}

type ResponseParameters struct {
	// Describes why a request was unsuccessful.
	migrate_to_chat_id *int64 //Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in int64erpreting it. But it has at most 52 significant bits, so a signed 64-bit int64eger or double-precision float type are safe for storing this identifier.
	retry_after        *int64 //Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

type InputMedia struct {
	// This object represents the content of a media message to be sent. It should be one of
	//     InputMediaAnimation
	//     InputMediaDocument
	//     InputMediaAudio
	//     InputMediaPhoto
	//     InputMediaVideo
	// InputMediaPhoto
	Type             string           //Type of the result, must be ["photo"|"video"|"animation"|"audio"|"document"
	media            string           //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the int64ernet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	caption          *string          //Optional. Caption of the photo/video/animation/audio/document to be sent, 0-1024 characters after entities parsing
	parse_mode       *string          //Optional. Mode for parsing entities in the photo/video/animation/audio/document caption. See formatting options for more details.
	caption_entities *[]MessageEntity //Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	has_spoiler      *bool            //Optional. Pass True if the photo/video/animation needs to be covered with a spoiler animation
	// InputMediaVideo

	// Represents a video to be sent.
	thumbnail          *string //Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	width              *int64  //Optional. Video/Animation width
	height             *int64  //Optional. Video/Animation height
	duration           *int64  //Optional. Video/Animation/audio duration in seconds
	supports_streaming *bool   //Optional. Pass True if the uploaded video is suitable for streaming
	// InputMediaAudio

	performer *string //Optional. Performer of the audio
	title     *string //Optional. Title of the audio
	// InputMediaDocument

	disable_content_type_detection *bool //Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
}

/*
string
// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
Sending files

There are three ways to send files (photos, stickers, audio, media, etc.):

    If the file is already stored somewhere on the Telegram servers, you don't need to reupload it: each file object has a file_id field, simply pass this file_id as a parameter instead of uploading. There are no limits for files sent this way.
    Provide Telegram with an HTTP URL for the file to be sent. Telegram will download and send the file. 5 MB max size for photos and 20 MB max for other types of content.
    Post the file using multipart/form-data in the usual way that files are uploaded via the browser. 10 MB max size for photos, 50 MB for other files.

Sending by file_id

    It is not possible to change the file type when resending by file_id. I.e. a video can't be sent as a photo, a photo can't be sent as a document, etc.
    It is not possible to resend thumbnails.
    Resending a photo by file_id will send all of its sizes.
    file_id is unique for each individual bot and can't be transferred from one bot to another.
    file_id uniquely identifies a file, but a file can have different valid file_ids even for the same bot.

Sending by URL

    When sending by URL the target file must have the correct MIME type (e.g., audio/mpeg for sendAudio, etc.).
    In sendDocument, sending by URL will currently only work for GIF, PDF and ZIP files.
    To use sendVoice, the file must have the type audio/ogg and be no more than 1MB in size. 1-20MB voice notes will be sent as files.
    Other configurations may work but we can't guarantee that they will.

Accent colors

Colors with identifiers 0 (red), 1 (orange), 2 (purple/violet), 3 (green), 4 (cyan), 5 (blue), 6 (pink) can be customized by app themes. Additionally, the following colors in RGB format are currently in use.

Color identifier	Light colors	Dark colors
7	E15052 F9AE63	FF9380 992F37
8	E0802B FAC534	ECB04E C35714
9	A05FF3 F48FFF	C697FF 5E31C8
10	27A910 A7DC57	A7EB6E 167E2D
11	27ACCE 82E8D6	40D8D0 045C7F
12	3391D4 7DD3F0	52BFFF 0B5494
13	DD4371 FFBE9F	FF86A6 8E366E
14	247BED F04856 FFFFFF	3FA2FE E5424F FFFFFF
15	D67722 1EA011 FFFFFF	FF905E 32A527 FFFFFF
16	179E42 E84A3F FFFFFF	66D364 D5444F FFFFFF
17	2894AF 6FC456 FFFFFF	22BCE2 3DA240 FFFFFF
18	0C9AB3 FFAD95 FFE6B5	22BCE2 FF9778 FFDA6B
19	7757D6 F79610 FFDE8E	9791FF F2731D FFDB59
20	1585CF F2AB1D FFFFFF	3DA6EB EEA51D FFFFFF

Profile accent colors

Currently, the following colors in RGB format are in use for profile backgrounds.

Color identifier	Light colors	Dark colors
0	BA5650	9C4540
1	C27C3E	945E2C
2	956AC8	715099
3	49A355	33713B
4	3E97AD	387E87
5	5A8FBB	477194
6	B85378	944763
7	7F8B95	435261
8	C9565D D97C57	994343 AC583E
9	CF7244 CC9433	8F552F A17232
10	9662D4 B966B6	634691 9250A2
11	3D9755 89A650	296A43 5F8F44
12	3D95BA 50AD98	306C7C 3E987E
13	538BC2 4DA8BD	38618C 458BA1
14	B04F74 D1666D	884160 A65259
15	637482 7B8A97	53606E 384654
}

Inline mode objects

Objects and methods used in the inline mode are described in the Inline mode section.
*/
