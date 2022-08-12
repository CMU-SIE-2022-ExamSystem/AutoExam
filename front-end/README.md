# exam-server/front-end

## Code Structure

```
front-end/
├─ public/
│  ├─ tinymce/                   /* The library of TinyMCE, an HTML editor */               
├─ src/
│  ├─ components/                /* Some reusable components such as layout */
│  │  ├─ questionTemplate/       /* Components that display a question */
│  │  │  ├─ readonly/            /* Question components, but without local storage */
│  │  ├─ AppLayout.tsx           /* The layout of all route components that provide access to Global State */
│  │  ├─ CountdownTimer.tsx      /* Timer used when students' are taking exams */
│  │  ├─ ErrorLayout.tsx         /* An alert component that pops out when error happens */
│  │  ├─ globalAlert.d.ts        /* The alert properties */
│  │  ├─ GlobalStateProvider.tsx /* Global States (states that persist over components) */
│  │  ├─ HTMLEditor.tsx
│  │  ├─ Question.tsx            /* Question layout */
│  │  ├─ RightBottomAlert.tsx    /* A component only used in ExamConfig, as a testing feature of global alert */
│  │  ├─ TopNavbar.tsx           /* The navbar at the top of most pages */
│  ├─ images/                    /* Image resources (such as Autolab banner) */
│  ├─ routes/                    /* Logics for each page */
│  │  ├─ auth/                   /* Authorization pages (Autolab OAuth) */
│  │  ├─ course/                 /* Course pages */
│  │  │  ├─ baseCourse/          /* Base course management (course numbers) */
│  │  │  ├─ config/              /* Create new exam configurations */
│  │  │  ├─ exams/               /* Students taking exams */
│  │  │  ├─ questionBanks/       /* Instructor add new questions and graders */
│  │  │  ├─ results/             /* Students check the feedback of exams */
│  │  │  ├─ Assessments.tsx      /* Show the list of exams and entry to other features */
│  │  │  ├─ Dashboard.tsx        /* Show the list of courses binded to Autolab */
│  │  ├─ Index.tsx               /* The welcoming page (the index page at root) */
│  ├─ utils/                     
│  ├─ App.css                    
│  ├─ App.tsx                    /* Page at Root */
│  ├─ index.css                  
│  ├─ index.tsx                  /* !! Global Entry !! */
├─ README.md                     
├─ package.json                  /* NPM document */
├─ tsconfig.json                 
├─ .env                          /* Environment Files */

```

This application uses *Node.js*. Please make sure the *Node.js* version in your environment satisfies ` version >= 14`. You may check your node version by running `node --version` in your terminal.

The global entry of this application is at `src/index.tsx`.

## .env file

The `.env` file should contain these variables:

- `NODE_DEV`: set to `development` for dev environment, and `production` in production environment.
- `REACT_APP_AUTOLAB_LOCATION`: The *Autolab* server host, for OAuth2.0 authentication purposes.
- `REACT_APP_BACKEND_API_ROOT`: The back end server host.



The following is the original Create React App README.

## Create React App 

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

### Available Scripts

In the project directory, you can run:

#### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

#### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

