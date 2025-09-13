package types

// This object represents a custom keyboard with reply options
// (see Introduction to bots (https://core.telegram.org/bots/features#keyboards)
// for details and examples). Not supported in channels and for messages sent
// on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"keyboard"`

	// Optional. Requests clients to always show the keyboard when the
	// regular keyboard is hidden. Defaults to false, in which case the
	// custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent bool `json:"is_persistent,omitempty"`

	// Optional. Requests clients to resize the keyboard vertically
	// for optimal fit (e.g., make the keyboard smaller if there are
	// just two rows of buttons). Defaults to false, in which case the
	// \custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// Optional. Requests clients to hide the keyboard as
	// soon as it's been used. The keyboard will still be
	// available, but clients will automatically display
	// the usual letter-keyboard in the chat - the user can
	// press a special button in the input field to see the
	// custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// Optional. The placeholder to be shown in the input
	// field when the keyboard is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Optional. Use this parameter if you want to show the keyboard to specific
	// users only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply to a message in the
	// same chat and forum topic, sender of the original message.
	//
	// Example: A user requests to change the bot's language, bot
	// replies to the request with a keyboard to select the new
	// language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// Upon receiving a message with this object, Telegram
// clients will remove the current custom keyboard and
// display the default letter-keyboard. By default,
// custom keyboards are displayed until a new keyboard
// is sent by a bot. An exception is made for one-time
// keyboards that are hidden immediately after the user
// presses a button (see ReplyKeyboardMarkup).
// Not supported in channels and for messages sent
// on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom
	// keyboard (user will not be able to summon
	// this keyboard; if you want to hide the
	// keyboard from sight but keep it accessible,
	// use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Optional. Use this parameter if you want to remove
	// the keyboard for specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message
	// object; 2) if the bot's message is a reply to a
	// message in the same chat and forum topic,
	// sender of the original message.
	//
	// Example: A user votes in a poll, bot returns
	// confirmation message in reply to the vote and
	// removes the keyboard for that user, while still
	// showing the keyboard with poll options to
	// users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// This object represents one button of the reply keyboard.
// At most one of the optional fields must be used to specify
// type of the button. For simple text buttons, String can be
// used instead of this object to specify the button text.
//
// Note: request_users and request_chat options will only
// work in Telegram versions released after 3 February, 2023.
// Older clients will display unsupported message.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are
	// used, it will be sent as a message when the button is pressed
	Text string `json:"text"`

	// Optional. If specified, pressing the button will open
	// a list of suitable users. Identifiers of selected
	// users will be sent to the bot in a “users_shared”
	// service message. Available in private chats only.
	RequestUsers any `json:"request_users,omitempty"`

	// Optional. If specified, pressing the button will
	// open a list of suitable chats. Tapping on a chat
	// will send its identifier to the bot in a “chat_shared”
	// service message. Available in private chats only.
	RequestChat any `json:"request_chat,omitempty"`

	// Optional. If True, the user's phone number will
	// be sent as a contact when the button is pressed.
	// Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`

	// Optional. If True, the user's current location
	// will be sent when the button is pressed. Available
	// in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// Optional. If specified, the user will be asked
	// to create a poll and send it to the bot when the
	// button is pressed. Available in private chats only.
	RequestPoll any `json:"request_poll,omitempty"`

	// Optional. If specified, the described Web App
	// will be launched when the button is pressed.
	// The Web App will be able to send a “web_app_data”
	// service message. Available in private chats only.
	WebApp any `json:"web_app,omitempty"`
}

// This object represents an inline keyboard
// (https://core.telegram.org/bots/features#inline-keyboards)
// that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an
	// Array of InlineKeyboardButton objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// This object represents one button of an inline keyboard.
// Exactly one of the optional fields must be used to specify
// type of the button.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`

	// Optional. HTTP or tg:// URL to be opened when
	// the button is pressed. Links tg://user?id=<user_id>
	// can be used to mention a user by their identifier
	// without using a username, if this is allowed by
	// their privacy settings.
	Url string `json:"url,omitempty"`

	// Optional. Data to be sent in a callback query
	// (https://core.telegram.org/bots/api#callbackquery) to
	// the bot when the button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// Optional. Description of the Web App (https://core.telegram.org/bots/webapps)
	// that will be launched when the user presses the button. The Web App
	// will be able to send an arbitrary message on behalf of the user
	// using the method answerWebAppQuery (https://core.telegram.org/bots/api#answerwebappquery).
	// Available only in private chats between a user and the bot. Not
	// supported for messages sent on behalf of a Telegram Business account.
	WebApp any `json:"web_app,omitempty"`

	// Optional. An HTTPS URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	// (https://core.telegram.org/widgets/login)
	LoginUrl any `json:"login_url,omitempty"`

	// Optional. If set, pressing the button will prompt the user to
	// select one of their chats, open that chat and insert the bot's
	// username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be
	// inserted. Not supported for messages sent in channel direct
	// messages chats and on behalf of a Telegram Business account.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. May be empty,
	// in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode in the
	// same chat - good for selecting something from multiple options. Not supported
	// in channels and for messages sent in channel direct messages chats and on
	// behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// Optional. If set, pressing the button will prompt the user to select one of
	// their chats of the specified type, open that chat and insert the bot's
	// username and the specified inline query in the input field. Not supported
	// for messages sent in channel direct messages chats and on behalf of a
	// Telegram Business account.
	SwitchInlineQueryChosenChat any `json:"switch_inline_query_chosen_chat,omitempty"`

	// Optional. Description of the button that copies
	// the specified text to the clipboard.
	CopyText any `json:"copy_text,omitempty"`

	// Optional. Description of the game that will be
	// launched when the user presses the button.
	//
	// NOTE: This type of button must always be the
	// first button in the first row.
	CallbackGame any `json:"callback_game,omitempty"`

	// Optional. Specify True, to send a Pay button. Substrings “⭐” and
	// “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	//
	// NOTE: This type of button must always be the first button in the
	// first row and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

// Upon receiving a message with this object, Telegram
// clients will display a reply interface to the user
// (act as if the user has selected the bot's message
// and tapped 'Reply'). This can be extremely useful
// if you want to create user-friendly step-by-step
// interfaces without having to sacrifice privacy mode
// (https://core.telegram.org/bots/features#privacy-mode).
// Not supported in channels and for messages sent
// on behalf of a Telegram Business account.
//
// Example: A poll bot for groups runs in privacy mode (only
// receives commands, replies to its messages and mentions).
// There could be two ways to create a new poll:
//
// - Explain the user how to send a command with parameters
// (e.g. /newpoll question answer1 answer2). May be appealing for
// hardcore users but lacks modern day polish.
//
// - Guide the user through a step-by-step process. 'Please
// send me your question', 'Cool, now let's add the first answer
// option', 'Great. Keep adding answer options, then send /done when you're ready'.
//
// The last option is definitely more attractive. And if you use
// ForceReply in your bot's questions, it will receive the user's
// answers even if it only receives replies, commands and mentions -
// without any extra work for the user.
type ForceReply struct {
	// Shows reply interface to the user,
	// as if they manually selected the
	// bot's message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`

	// Optional. The placeholder to be shown in the input
	// field when the reply is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Optional. Use this parameter if you want to
	// force reply from specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message
	// object; 2) if the bot's message is a reply to a
	// message in the same chat and forum topic,
	// sender of the original message.
	Selective bool `json:"selective,omitempty"`
}
