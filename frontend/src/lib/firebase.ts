import { initializeApp, getApps } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "AIzaSyCI6Sny25W-XtA8xIp1YV13epZkUxZGe4s",
  authDomain: "term6-hiroto-uchida.firebaseapp.com",
  databaseURL: "https://term6-hiroto-uchida-default-rtdb.firebaseio.com",
  projectId: "term6-hiroto-uchida",
  storageBucket: "term6-hiroto-uchida.firebasestorage.app",
  messagingSenderId: "241499864821",
  appId: "1:241499864821:web:c39f3db60824b5b5df8fbc"
};

// Initialize Firebase only if it hasn't been initialized already
const app = !getApps().length ? initializeApp(firebaseConfig) : getApps()[0];
console.log(app)

// Log any initialization errors
if (!app) {
  console.error('Firebase initialization error');
}

export const auth = getAuth(app);

