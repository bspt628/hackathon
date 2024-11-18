import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

// Firebase の初期化
const auth = getAuth();

signInWithEmailAndPassword(auth, "test9@example.com", "password")
  .then((userCredential) => {
    const user = userCredential.user;
    user.getIdToken().then((idToken) => {
      console.log("ID Token:", idToken);

      // ID トークンを画面に表示
        const tokenElement = document.getElementById("id-token");
        // メッセージを表示
        
      tokenElement.textContent = `ID Token: ${idToken}`;
    });
  })
  .catch((error) => {
    console.error("Error:", error);
  });
