package db

var firstNames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe", "adam", "ben", "chris", "daniel",
	"edward", "fred", "greg", "harry", "ian", "james", "kevin", "liam", "michael"}

var lastNames = []string{
	"smith", "johnson", "williams", "jones", "brown", "davis", "miller", "wilson",
	"moore", "taylor", "anderson", "thomas", "jackson", "white", "harris", "martin",
	"thompson", "garcia", "martinez", "robinson", "clark", "rodriguez", "lewis", "lee",
	"walker", "hall", "allen", "young", "hernandez", "king", "wright", "lopez",
	"hill", "scott", "green", "adams", "baker", "gonzalez", "nelson", "carter",
	"mitchell", "perez", "roberts", "turner", "phillips", "campbell", "parker", "evans",
	"edwards", "collins", "stewart", "sanchez", "morris", "rogers", "reed", "cook",
	"morgan", "bell", "murphy", "bailey", "rivera", "cooper", "richardson", "cox",
}

var emails = []string{
	"gmail.com", "yahoo.com", "outlook.com", "icloud.com", "aol.com", "protonmail.com", "zoho.com", "yandex.com",
}

var titles = []string{
	"Software Engineer", "Data Scientist", "Product Manager", "UX Designer", "UI Designer", "Frontend Developer", "Backend Developer", "Fullstack Developer",
	"DevOps Engineer", "Site Reliability Engineer", "Security Engineer", "Network Engineer", "Machine Learning Engineer", "AI Engineer", "Cloud Engineer", "Database Administrator",
	"Go", "Python", "Java", "JavaScript", "Ruby", "PHP", "C++", "C#", ".NET", "Swift", "Kotlin", "Rust", "Scala", "Perl", "Spring",
	"Postgres", "MySQL", "SQLite", "MongoDB", "Redis", "Cassandra", "DynamoDB", "MariaDB", "Oracle", "MSSQL", "CockroachDB", "Couchbase", "Elasticsearch", "Firebase", "Aurora",
	"AWS", "Azure", "GCP", "DigitalOcean", "Heroku", "Netlify", "Vercel", "Cloudflare", "Fastly", "Akamai", "CloudFront", "CloudWatch", "CloudTrail", "CloudFormation", "CloudSearch",
	"React", "Vue", "Angular", "Svelte", "Ember", "Backbone", "jQuery", "D3", "Three", "Chart", "Bootstrap", "Tailwind", "Material", "Ant", "Chakra",
}

func SeedDatabase(db PostgresDb) error {

	return nil
}
