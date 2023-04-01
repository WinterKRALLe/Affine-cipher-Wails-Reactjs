import { useState } from 'react';
import { Encrypt, Decrypt } from "../wailsjs/go/main/App";

import { ReactComponent as Key } from "./assets/key.svg"
import "./style.scss"

function App() {
    const [keyA, setKeyA] = useState(0)
    const [keyB, setKeyB] = useState(0)
    const [inputEncrypt, setInputEncrypt] = useState("");
    const [outputEncrypt, setOutputEncrypt] = useState("");
    const [inputDecrypt, setInputDecrypt] = useState("");
    const [outputDecrypt, setOutputDecrypt] = useState("");

    const updateKeyA = (e) => setKeyA(e.target.value)
    const updateKeyB = (e) => setKeyB(e.target.value)

    const updateInputEncrypt = (e) => setInputEncrypt(e.target.value)
    const updateOutputEncrypt = (text) => setOutputEncrypt(text)

    const updateInputDecrypt = (e) => setInputDecrypt(e.target.value)
    const updateOutputDecrypt = (text) => setOutputDecrypt(text)

    function encrypt() {
        Encrypt(inputEncrypt, keyA, keyB).then(updateOutputEncrypt)
    }
    function decrypt() {
        Decrypt(inputDecrypt, keyA, keyB).then(updateOutputDecrypt)
    }



    return (
        <div className="App">
            <div className="keys">
                <div className="keyBox">
                    <label htmlFor="klicA">Key A</label>
                    <div className="inputBox">
                        <Key />
                        <input type="text" name="keyA" id="keyA" onChange={updateKeyA} />
                    </div>
                </div>
                <div className="keyBox">
                    <label htmlFor="klicA">Key B</label>
                    <div className="inputBox">
                        <Key />
                        <input type="text" name="keyB" id="keyB" onChange={updateKeyB} />
                    </div>
                </div>
            </div>

            <div className="inputs">
                <div className="keyBox">
                    <div>
                        <label htmlFor="encrypt">Enter text: </label>
                        <div className="inputBox">
                            <input type="text" name="inputEncrypt" id="encrypt" onChange={updateInputEncrypt} />
                            <button onClick={encrypt}>Encrypt</button>
                        </div>
                    </div>
                    <div>
                        <label htmlFor="outputEncrypt">Output</label>
                        <div className="inputBox">
                            <textarea name="outputEncrypt" id="outputEncrypt" readOnly value={outputEncrypt} />
                        </div>
                    </div>
                </div>
                <div className="keyBox">
                    <div>
                        <label htmlFor="Decrypt">Enter text: </label>
                        <div className="inputBox">
                            <input type="text" name="inputEncrypt" id="Decrypt" onChange={updateInputDecrypt} />
                            <button onClick={decrypt}>Decrypt</button>
                        </div>
                    </div>
                    <div>
                        <label htmlFor="outputEncrypt">Output</label>
                        <div className="inputBox">
                            <textarea name="outputDecrypt" id="outputDecrypt" readOnly value={outputDecrypt} />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default App
