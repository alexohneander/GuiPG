import { useState } from 'react';
import logo from './assets/images/gopenpgp.svg';
import './App.css';
import { Encrypt, Sign, Decrypt } from "../wailsjs/go/main/App";

function App() {
  const [textContent, setTextContent] = useState("");
  const updateTextContent = (e: any) => setTextContent(e.target.value);
  const [privateKey, setPrivateKey] = useState('');
  const updatePrivateKey = (e: any) => setPrivateKey(e.target.value);
  const [passphrase, setPassphrase] = useState('');
  const updatePassphrase = (e: any) => setPassphrase(e.target.value);

  function encrypt() {
    // Verwende aktuellen textContent als Input für Verschlüsselung
    Encrypt(textContent, privateKey).then((result) => {
      setTextContent(result); // Ergebnis ins selbe Feld
    });
  }

  function sign() {
    Sign(textContent, privateKey, passphrase).then((result) => {
      setTextContent(result);
    });
  }

  function decrypt() {
    Decrypt(textContent, privateKey, passphrase).then((result) => {
      setTextContent(result);
    });
  }

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div className="result">Please enter your Private/Public -Key below 👇</div>
      <div className="input-box">
        <textarea
          id="keybox"
          onChange={updatePrivateKey}
          className="txtarea"
          placeholder="Public/Private Key"
          value={privateKey}
        />
      </div>
      <div className="input-box">
        <label>Password </label>
        <input
          id="passphrase"
          onChange={updatePassphrase}
          type="password"
          placeholder="Password"
          value={passphrase}
        />
      </div>
      <div id="result" className="result">Message / Result 👇</div>
      <div className="input-box">
        <textarea
          id="msgbox"
          className="txtarea"
          onChange={updateTextContent}
          placeholder="Message"
          value={textContent}
        />
      </div>
      <div id="input" className="input-box">
        <button className="btn" onClick={encrypt}>Encrypt</button>
        <button className="btn" onClick={sign}>Sign</button>
        <button className="btn" onClick={decrypt}>Decrypt</button>
      </div>
    </div>
  )
}

export default App
