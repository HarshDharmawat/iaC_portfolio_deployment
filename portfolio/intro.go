package main

type Intro struct {
	Name        string            `json:"name"`
	Role        string            `json:"role"`
	Description string            `json:"description"`
	Skills      []string          `json:"skills"`
	Socials     map[string]string `json:"socials"`
}

var myIntro = Intro{
	Name:        "Harsh Dharmawat",
	Role:        "Cloud & DevOps Engineer",
	Description: "Cloud Project Trainee at Data Oceano | Cloud & DevOps | AWS | Docker | CI/CD | Linux | Kubernetes | IaC | Cloud-Native Automation & Architecture | Data Structures & Algorithms | MERN",
	Skills: []string{
			"C++", "Go", "Linux", "Networking",
			"Docker", "Kubernetes", "CI/CD",
			"AWS", "Terraform", "Ansible",
			"Cloud-Native Architecture",
		},
	Socials: map[string]string{
			"X":          "https://x.com/DharmawatHarsh",
			"LinkedIn":   "https://www.linkedin.com/in/harsh-dharmawat/",
			"GitHub":     "https://github.com/HarshDharmawat",
			"LeetCode":   "https://leetcode.com/u/harshdharmawat/",
			"CodeChef":   "https://www.codechef.com/users/harshdharmawat",
			"CodeForces": "https://codeforces.com/profile/harshdharmawat",
			"Medium":     "https://medium.com/@harshdharmawat",
			"Codolio":	  "https://codolio.com/profile/harshdharmawat_10",
	},
}