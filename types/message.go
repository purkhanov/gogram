package types

type Message struct {
	// Unique message identifier inside this chat.
	// In specific instances (e.g., message containing
	// a video sent to a big chat), the server might
	// automatically schedule a message instead of
	// sending it immediately. In such cases, this
	// field will be 0 and the relevant message will
	// be unusable until it is actually sent
	MessageID uint `json:"message_id"`

	// Optional. Unique identifier of a message thread
	// to which the message belongs; for supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Optional. Information about the direct messages chat topic that contains the message
	DirectMessagesTopic *DirectMessagesTopic `json:"direct_messages_topic,omitempty"`

	// Optional. Sender of the message; may be empty for
	// messages sent to channels. For backward compatibility,
	// if the message was sent on behalf of a chat, the
	// field contains a fake sender user in non-channel chats
	From *User `json:"from,omitempty"`

	// Optional. Sender of the message when sent on behalf
	// of a chat. For example, the supergroup itself for
	// messages sent by its anonymous administrators or a
	// linked channel for messages automatically forwarded
	// to the channel's discussion group. For backward
	// compatibility, if the message was sent on behalf
	// of a chat, the field from contains a fake sender
	// user in non-channel chats.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Optional. If the sender of the message boosted the
	// chat, the number of boosts added by the user
	SenderBoostCount int `json:"sender_boost_count,omitempty"`

	// Optional. The bot that actually sent the message on behalf
	// of the business account. Available only for outgoing
	// messages sent on behalf of the connected business account.
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`

	// Date the message was sent in Unix time. It is always a
	// positive number, representing a valid date.
	Date int `json:"date"`

	// Optional. Unique identifier of the business connection
	// from which the message was received. If non-empty,
	// the message belongs to a chat of the corresponding business
	// account that is independent from any potential bot chat
	// which might share the same identifier.
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`

	// Chat the message belongs to
	Chat *Chat `json:"chat"`

	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`

	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// Optional. True, if the message is a channel post that was
	// automatically forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// Optional. For replies in the same chat and message
	// thread, the original message. Note that the Message
	// object in this field will not contain further
	// reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// Optional. Information about the message that is being
	// replied to, which may come from another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`

	// Optional. For replies that quote part of the
	// original message, the quoted part of the message
	Qoute *TextQuote `json:"quote,omitempty"`

	// Optional. For replies to a story, the original story
	ReplyToStory *Story `json:"reply_to_story,omitempty"`

	// Optional. Identifier of the specific checklist task that is being replied to
	ReplyToChecklistTaskID int `json:"reply_to_checklist_task_id,omitempty"`

	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`

	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// Optional. True, if the message was sent by an
	// implicit action, for example, as an away or a
	// greeting business message, or as a scheduled message
	IsFromOffline bool `json:"is_from_offline,omitempty"`

	// Optional. The unique identifier of a media
	// message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// Optional. Signature of the post author for
	// messages in channels, or the custom title
	// of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// Optional. The number of Telegram Stars that were
	// paid by the sender of the message to send it
	PaidStarCount int `json:"paid_star_count,omitempty"`

	// Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text,omitempty"`

	// Optional. For text messages, special entities like usernames,
	// URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// Optional. Options used for link preview generation for the message,
	// if it is a text message and link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Optional. Information about suggested post parameters if
	// the message is a suggested post in a channel direct messages
	// chat. If the message is an approved or declined suggested
	// post, then it can't be edited.
	SuggestedPostInfo *SuggestedPostInfo `json:"suggested_post_info,omitempty"`

	// Optional. Unique identifier of the message effect added to the message
	EffectID string `json:"effect_id,omitempty"`

	// Optional. Message is an animation, information about
	// the animation. For backward compatibility, when this
	// field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message contains paid media; information about the paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. Caption for the animation, audio,
	// document, paid media, photo, video or voice
	Caption string `json:"caption,omitempty"`

	// Optional. For messages with a caption, special entities like
	// usernames, URLs, bot commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a checklist
	Checklist *Checklist `json:"checklist,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Optional. Message is a game, information about the game.
	Game *Game `json:"game,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a venue, information about
	// the venue. For backward compatibility, when this
	// field is set, the location field will also be set
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. New members that were added to the group
	// or supergroup and information about them (the bot
	// itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// Optional. A member was removed from the group, information
	// about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// Optional. Service message: the supergroup has been
	// created. This field can't be received in a message
	// coming through updates, because bot can't be a member
	// of a supergroup when it is created. It can only be
	// found in reply_to_message if someone replies to a very
	// first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// Optional. Service message: the channel has been created.
	// This field can't be received in a message coming through
	// updates, because bot can't be a member of a channel when
	// it is created. It can only be found in reply_to_message
	// if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	// Optional. The group has been migrated to a supergroup
	// with the specified identifier. This number may have
	// more than 32 significant bits and some programming
	// languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`

	// Optional. The supergroup has been migrated from a group
	// with the specified identifier. This number may have more
	// \than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed
	// 64-bit integer or double-precision float type are safe
	// for storing this identifier.
	MigrateFromChatID int `json:"migrate_from_chat_id,omitempty"`

	// Optional. Specified message was pinned. Note that the
	// Message object in this field will not contain further
	// reply_to_message fields even if it itself is a reply.
	PinnedMessage *MaybeInaccessibleMessage `json:"pinned_message,omitempty"`

	// Optional. Message is an invoice for a payment, information
	// about the invoice. https://core.telegram.org/bots/api#payments
	Invoice *Invoice `json:"invoice,omitempty"`

	// Optional. Message is a service message about a
	// successful payment, information about the payment.
	// https://core.telegram.org/bots/api#payments
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// Optional. Message is a service message about a refunded
	// payment, information about the payment.
	// https://core.telegram.org/bots/api#payments
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`

	// Optional. Service message: users were shared with the bot
	UsersShared *UsersShared `json:"users_shared,omitempty"`

	// Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared,omitempty"`

	// Optional. Service message: a regular gift was sent or received
	Gift *GiftInfo `json:"gift,omitempty"`

	// Optional. Service message: a unique gift was sent or received
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`

	// Optional. The domain name of the website on which the user has logged in.
	// https://core.telegram.org/widgets/login
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// Optional. Service message: the user allowed the bot to write
	// messages after adding it to the attachment or side menu,
	// launching a Web App from a link, or accepting an explicit
	// request from a Web App sent by the method requestWriteAccess
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`

	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`

	// Optional. Service message. A user in the chat triggered
	// another user's proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// Optional. Service message: user boosted the chat
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`

	// Optional. Service message: chat background set
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`

	// Optional. Service message: some tasks in a checklist were marked as done or not done
	ChecklistTasksDone *ChecklistTasksDone `json:"checklist_tasks_done,omitempty"`

	// Optional. Service message: tasks were added to a checklist
	ChecklistTasksAdded *ChecklistTasksAdded `json:"checklist_tasks_added,omitempty"`

	// Optional. Service message: the price for paid messages in
	// the corresponding direct messages chat of a channel has changed
	DirectMessagePriceChanged *DirectMessagePriceChanged `json:"direct_message_price_changed,omitempty"`

	// Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`

	// Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`

	// Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`

	// Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`

	// Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`

	// Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`

	// Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`

	// Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`

	// Optional. Service message: the price for paid messages has changed in the chat
	PaidMessagePriceChanged *PaidMessagePriceChanged `json:"paid_message_price_changed,omitempty"`

	// Optional. Service message: a suggested post was approved
	SuggestedPostApproved *SuggestedPostApproved `json:"suggested_post_approved,omitempty"`

	// Optional. Service message: approval of a suggested post has failed
	SuggestedPostApprovalFailed *SuggestedPostApprovalFailed `json:"suggested_post_approval_failed,omitempty"`

	// Optional. Service message: a suggested post was declined
	SuggestedPostDeclined *SuggestedPostDeclined `json:"suggested_post_declined,omitempty"`

	// Optional. Service message: payment for a suggested post was received
	SuggestedPostPaid *SuggestedPostPaid `json:"suggested_post_paid,omitempty"`

	// Optional. Service message: payment for a suggested post was refunded
	SuggestedPostRefunded *SuggestedPostRefunded `json:"suggested_post_refunded,omitempty"`

	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`

	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`

	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`

	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`

	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// Optional. Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// This object represents a unique message identifier.
type MessageId struct {
	// Unique message identifier. In specific instances
	// (e.g., message containing a video sent to a big chat),
	// the server might automatically schedule a message
	// instead of sending it immediately. In such cases,
	// this field will be 0 and the relevant message will
	// be unusable until it is actually sent
	MessageID int `json:"message_id"`
}

type InaccessibleMessage struct {
	// Chat the message belonged to
	Chat *Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int `json:"message_id"`

	// Always 0. The field can be used to differentiate
	// regular and inaccessible messages.
	Date int `json:"date"`
}

// This object describes a message that can be
// inaccessible to the bot. It can be one of
type MaybeInaccessibleMessage struct {
	*Message
	*InaccessibleMessage
}

type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (@username),
	// “hashtag” (#hashtag or #hashtag@chatusername), “cashtag”
	// ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot),
	// “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic”
	// (italic text), “underline” (underlined text), “strikethrough”
	// (strikethrough text), “spoiler” (spoiler message), “blockquote”
	// (block quotation), “expandable_blockquote” (collapsed-by-default
	// block quotation), “code” (monowidth string), “pre” (monowidth block),
	// “text_link” (for clickable text URLs), “text_mention” (for users
	// without usernames), “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`

	// Length of the entity in UTF-16 code units
	Length int `json:"length"`

	// Optional. For “text_link” only, URL that will be
	// opened after user taps on the text
	Url string `json:"url,omitempty"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`

	// Optional. For “custom_emoji” only, unique identifier of the custom emoji.
	// Use getCustomEmojiStickers to get full information about the sticker
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type MessageOrigin struct{}

type DirectMessagesTopic struct {
	// Unique identifier of the topic
	TopicID int `json:"topic_id"`

	// Optional. Information about the user that created the topic.
	// Currently, it is always present
	User User `json:"user,omitempty"`
}

// Describes a service message about a change in the
// price of direct messages sent to a channel chat.
type DirectMessagePriceChanged struct {
	// True, if direct messages are enabled for the channel chat; false otherwise
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`

	// Optional. The new number of Telegram Stars that
	// must be paid by users for each direct message sent
	// to the channel. Does not apply to users who have been
	// exempted by administrators. Defaults to 0.
	DirectMessageStarCount int `json:"direct_message_star_count,omitempty"`
}

// This object contains information about the quoted part
// of a message that is replied to by the given message.
type TextQuote struct {
	// Text of the quoted part of a message that
	// is replied to by the given message
	Text string `json:"text"`

	// Optional. Special entities that appear in the quote.
	// Currently, only bold, italic, underline, strikethrough,
	// spoiler, and custom_emoji entities are kept in quotes.
	Entities []MessageEntity `json:"entities,omitempty"`

	// Approximate quote position in the original message in
	// UTF-16 code units as specified by the sender
	Position int `json:"position,omitempty"`

	// Optional. True, if the quote was chosen manually
	// by the message sender. Otherwise, the quote was
	// added automatically by the server.
	IsManual bool `json:"is_manual,omitempty"`
}

// Describes a service message about a change
// in the price of paid messages within a chat.
type PaidMessagePriceChanged struct {
	// The new number of Telegram Stars that must
	// be paid by non-administrator users of the
	// supergroup chat for each sent message
	PaidMessageStarCount int `json:"paid_message_star_count"`
}
