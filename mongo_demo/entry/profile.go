package entry

type Profile struct {
	ID        string   `bson:"_id"`
	Name      string   `bson:"name"`
	Phone     string   `bson:"phone"`
	Email     string   `bson:"email"`
	Ipv4      string   `bson:"ipv4"`
	Timestamp int      `bson:"timestamp"`
	Hobby     []string `bson:"hobby"`
}

type ProfileInsert struct {
	Name      string   `bson:"name"`
	Phone     string   `bson:"phone"`
	Email     string   `bson:"email"`
	Ipv4      string   `bson:"ipv4"`
	Timestamp int      `bson:"timestamp"`
	Hobby     []string `bson:"hobby"`
}
