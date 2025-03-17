package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math/rand"
	"strings"
	"time"
)

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

func SeedCandidatesTable(db PostgresDb) error {
	startTime := time.Now()
	const batchSize = 1000
	const totalRecords = 1_000_000
	batches := totalRecords / batchSize

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for b := 0; b < batches; b++ {
		if b%100 == 0 {
			fmt.Printf("Processing batch %d of %d\n", b, batches)
		}
		pgxBatch := &pgx.Batch{}
		for i := 0; i < batchSize; i++ {
			firstName := firstNames[r.Intn(len(firstNames))]
			lastName := lastNames[r.Intn(len(lastNames))]
			email := fmt.Sprintf("%s.%s%d@%s", firstName, lastName, r.Intn(23457111), emails[r.Intn(len(emails))])

			skillsNum := r.Intn(5) + 1
			usedSkills := make(map[int]struct{})
			title := strings.Builder{}

			for j := 0; len(usedSkills) < skillsNum; j++ {
				rSkill := r.Intn(skillsNum)
				if _, ok := usedSkills[rSkill]; ok {
					continue
				}

				usedSkills[rSkill] = struct{}{}
				title.WriteString(titles[rSkill])
				title.WriteString(" | ")
			}

			titleStr := title.String()

			pgxBatch.Queue("INSERT INTO candidates (first_name, last_name, email, title) VALUES ($1, $2, $3, $4)", firstName, lastName, email, titleStr)
		}

		bRes := db.DB.SendBatch(context.Background(), pgxBatch)
		_, err := bRes.Exec()
		if err != nil {
			log.Fatal(err)
		}

		_ = bRes.Close()
		fmt.Printf("Batch %d of %d processed\n", b, batches)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(startTime))
	return nil
}

func SeedCommentsTable(db PostgresDb) error {
	const batchSize = 1000
	const totalRecords = 1_000_000
	batches := totalRecords / batchSize

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for b := 0; b < batches; b++ {
		if b%100 == 0 {
			fmt.Printf("Processing batch %d of %d\n", b, batches)
		}
		pgxBatch := &pgx.Batch{}
		for i := 0; i < batchSize; i++ {
			postID := r.Intn(100_000) + 1
			userID := r.Intn(1_000_000) + 1
			content := fmt.Sprintf("This is a comment by user %d on post %d", userID, postID)

			pgxBatch.Queue("INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3)", postID, userID, content)
		}

		_, err := db.DB.SendBatch(context.Background(), pgxBatch).Exec()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
