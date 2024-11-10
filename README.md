# LTI Applicant Tracking System

## Overview

LTI is an innovative Applicant Tracking System (ATS) designed specifically for medium-sized businesses. It leverages advanced AI data analysis to streamline the recruitment process, enhance HR efficiency, and foster real-time collaboration between recruiters and hiring managers.

## Project Structure

```
lti-ats/
├── backend/           # Go backend with DDD architecture
├── frontend/         # React frontend with Tailwind CSS
├── docs/            # Project documentation
├── migrations/      # Database migrations
└── docker-compose.yml
```

## Technologies

- **Backend**: Go 1.22.4, PostgreSQL
- **Frontend**: React 18, Tailwind CSS
- **Infrastructure**: Docker, Docker Compose
- **Documentation**: Markdown, C4 Model

## Quick Start

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/lti-ats.git
cd lti-ats
```

2. **Set up environment variables**
```bash
cp .env.example .env
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

3. **Start the database**
```bash
docker compose up -d
```

4. **Run the backend**
```bash
cd backend
go run main.go
```

5. **Run the frontend**
```bash
cd frontend
npm install
npm start
```

## Documentation

- [Project Context](docs/ProjectContext-RJD.md)
- User Stories
- Backend Documentation
- [Frontend Documentation](frontend/README.md)

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the `LICENSE.md` file for details.

## Version

See `VERSION` file for current version information.