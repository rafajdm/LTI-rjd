Here's the updated frontend README with more specific information about our LTI ATS project:

# LTI ATS Frontend

## Overview

Frontend application for the LTI Applicant Tracking System, implementing a modern UI for candidate profile management and recruitment workflow.

## Project Structure

```
frontend/
├── public/            # Static files
├── src/
│   ├── components/   # Reusable React components
│   │   ├── Layout/           # Application layout components
│   │   └── CandidateProfile/ # Candidate management components
│   ├── pages/       # Page components
│   │   └── HomePage.js
│   └── services/    # API services
├── package.json
└── tailwind.config.js
```

## Features

- Candidate Profile Management
- Modern UI with Tailwind CSS
- Responsive Layout
- Form Validation
- API Integration with Backend

## Quick Start

1. **Install Dependencies**
```bash
npm install
```

2. **Environment Setup**
```bash
cp .env.example .env
```

3. **Start Development Server**
```bash
npm start
```

## Available Scripts

- `npm start` - Run development server on port 3000
- `npm test` - Run test suite
- `npm run build` - Create production build

## Components

### Layout
- Sidebar navigation
- Content area
- Responsive design

### CandidateProfileForm
- Form validation
- API integration
- Success/Error handling

## Development

- React 18
- Tailwind CSS for styling
- Environment-based configuration
- Component-based architecture

## Testing
```bash
npm test
```

## Building
```bash
npm run build
```

## Contributing

See CONTRIBUTING.md in the root directory.

## Documentation

For detailed documentation about:
- [Project Context](../docs/ProjectContext-RJD.md)
- User Stories
- API Documentation