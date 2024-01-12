package api

/*
Telegram Passport is a unified authorization method
for services that require personal identification.
Users can upload their documents once, then instantly share their data
with services that require real-world ID (finance, ICOs, etc.).
Please see the manual for details.
*/

type PassportData struct {
	// Describes Telegram Passport data shared with the bot by the user.
	data        []EncryptedPassportElement //Array with information about documents and other Telegram Passport elements that was shared with the bot
	credentials EncryptedCredentials       //Encrypted credentials required to decrypt the data
}

type PassportFile struct {
	// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
	file_id        string //Identifier for this file, which can be used to download or reuse the file
	file_unique_id string //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	file_size      int    //File size in bytes
	file_date      int    //Unix time when the file was uploaded
}

type EncryptedPassportElement struct {
	// Describes documents or other Telegram Passport elements shared with the bot by the user.
	Type         string         //Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	data         string         //Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	phone_number string         //Optional. User's verified phone number, available only for “phone_number” type
	email        string         //Optional. User's verified email address, available only for “email” type
	files        []PassportFile //Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	front_side   PassportFile   //Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	reverse_side PassportFile   //Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	selfie       PassportFile   //Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	translation  []PassportFile //Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	hash         string         //Base64-encoded element hash for using in PassportElementErrorUnspecified
}
type EncryptedCredentials struct {
	// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
	data   string //Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	hash   string //Base64-encoded data hash for data authentication
	secret string //Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}
