
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

# AWS Infrastructure-as-Code & Go Deployment Pipeline

An enterprise-grade, zero-touch deployment architecture for a containerized Go REST API. This project utilizes **Infrastructure as Code (IaC)** to provision AWS resources and **Ansible** to orchestrate a fully automated CI/CD deployment pipeline.

## Architecture & Tech Stack

*   **Application:** Go (Golang) REST API
*   **Load Balancer:** Traefik (Reverse Proxy & API Gateway)
*   **Observability:** Prometheus & Grafana
*   **Containerization:** Docker (Multi-stage builds optimizing the final image to ~25MB)
*   **Infrastructure as Code (IaC):** AWS CloudFormation
*   **Configuration Management:** Ansible
*   **Cloud Provider (AWS):** EC2, Elastic Container Registry (ECR), IAM Profiles, Security Groups

## The Deployment Pipeline

The deployment process is broken down into three automated playbooks:

1.  **The Architect (`deploy-aws-infra.yml`):** Uses CloudFormation to provision the VPC, Security Groups, IAM Roles, ECR repository, and the EC2 host.
2.  **The Delivery (`build-and-push.yml`):** Compiles the Go application natively, builds the optimized Docker container, and securely authenticates & pushes the image to AWS ECR.
3.  **The Application (`deploy-app.yml`):** Connects to the EC2 instance via SSH, installs Docker/AWS CLI, authenticates with ECR using credential-less IAM roles, pulls the latest image, and restarts the container with zero manual intervention.
4.  **The Router (`deploy-lb.yml`):** Deploys Traefik on the Edge Load Balancer instance, generating dynamic configurations to securely route public traffic to the App Server.
5.  **The Watcher (`deploy-monitoring.yml`):** Provisions Prometheus and Grafana on the Observability instance, configuring it to scrape network metrics from Traefik over the internal network.

## Security Best Practices Implemented
*   **Network Segmentation:** The Go API is strictly sandboxed. AWS Security Groups are configured so the App server drops all direct public traffic, only trusting the internal IP of the Traefik Load Balancer.
*   **VPC-Internal Scraping:** Prometheus scrapes Traefik metrics using internal AWS 172.31.x.x IPs. The metrics port is firewalled from the public internet, ensuring traffic data cannot be intercepted externally.
*   **Credential-less Server Access:** The EC2 instance securely pulls Docker images from ECR using an attached IAM Instance Profile, eliminating the need for hardcoded AWS credentials on the server.
*   **Strict Secrets Management:** `hosts.ini`, AWS `.pem` keys, and Ansible `all.yml` variables are strictly excluded via `.gitignore` to prevent secret leakage.
*   **Optimized Attack Surface:** The Docker image uses multi-stage builds to strip away OS vulnerabilities and build tools, leaving only the compiled Go binary in the final production container.

## Project Structure

```text
├── cloudformation/
│   └── core-infrastructure.yml    # AWS resource definitions
|   └── loadbalancer.yml           # Traefik load balancer definition
|   └── observabilityinfra.yml     # Prometheus and Grafana defition
├── group_vars/
│   └── all.example.yml            # Dummy variable templates
├── inventory/
│   └── hosts.example.ini          # Dummy server inventory
├── playbooks/                     
│   ├── deploy-aws-infra.yml       # Provisions AWS infrastructure
│   ├── build-and-push.yml         # Builds & pushes Docker image to ECR
│   └── deploy-app.yml             # Pulls and runs image on EC2
|   └── deploy-lb.yml              # Deploys load balancer on an EC2 instance
|   └── deploy-montoring.yml       # Deploys Promethus and Grafana on an EC2 instance
└── portfolio/                     
    ├── Dockerfile                 # Multi-stage container instructions
    └── main.go                    # Application entry point
    └── handlers.go
    └── routes.go
    └── intro.go
