# exam-server/front-end

## Code Structure

```
front-end/
├─ public/
├─ src/
│  ├─ components/				/* Some reusable components such as layout */
│  ├─ images/
│  ├─ routes/                   /* Logics for each page */
│  │  ├─ auth/                  /* Authorization pages (Autolab OAuth) */
│  │  ├─ course/                /* Course pages */
│  ├─ utils/
│  ├─ App.css
│  ├─ App.tsx                   /* Page at Root */
│  ├─ index.css
│  ├─ index.tsx                 /* !! Global Entry !! */
├─ README.md
├─ package.json                 /* NPM document */
├─ tsconfig.json
├─ .env                         /* Environment Files */

```

This application uses *Node.js*. Please make sure the *Node.js* version in your environment satisfies ` version >= 14`. You may check your node version by running `node --version` in your terminal.

The global entry of this application is at `src/index.tsx`.

Application under development.



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

#### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

#### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

