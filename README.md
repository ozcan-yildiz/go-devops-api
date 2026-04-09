# Go DevOps Project: CI/CD Pipeline with AWS EC2 & Docker

This is a simple Go application that demonstrates a complete DevOps lifecycle. The goal was to build a classic, reliable CI/CD pipeline using industry-standard tools.

### Project Overview
The application is a basic Go web server, but the focus is on the infrastructure. I used a **container-first approach** to ensure the app runs the same everywhere.

### Tech Stack & Tools
* **Language:** Go (Golang)
* **Containerization:** Docker
* **Registry:** AWS ECR (Elastic Container Registry)
* **Compute:** AWS EC2 (Ubuntu Instance)
* **CI/CD:** GitHub Actions

### How the Pipeline Works
I designed the CI/CD flow to be fully automated:

1.  **Continuous Integration:** Every time I push code to the `main` branch, GitHub Actions triggers a workflow. It sets up the Go environment and runs unit tests to make sure everything is okay.
2.  **Containerization:** After successful tests, the app is Dockerized. The image is then pushed to **AWS ECR** with a unique tag.
3.  **Continuous Deployment:** Finally, the pipeline connects to my **AWS EC2** instance via SSH. It pulls the latest image from ECR, stops the old container, and starts the new one.

### Security & Configuration
Security was a key part of this setup. **All sensitive data like AWS keys and SSH credentials are securely managed via GitHub Secrets to ensure no private information is exposed in the source code.** Additionally, I implemented **IAM Roles** on the EC2 instance to grant it read-only access to ECR, following the **Principle of Least Privilege**.

### Key Features
* **Automated Testing:** No code goes to production without passing tests.
* **Secure Access:** Used IAM Roles for EC2-to-ECR communication instead of hardcoding credentials.
* **Least Privilege:** All AWS policies are restricted to only what is necessary.
* **Zero Manual Work:** From code push to live URL, everything is automated.
