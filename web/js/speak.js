document.getElementById("startRecording").addEventListener("click", initFunction);

const isRecording = document.getElementById("isRecording");

function initFunction() {
    // Display recording
    async function getUserMedia(constraints) {
        if (window.navigator.mediaDevices) {
            return window.navigator.mediaDevices.getUserMedia(constraints);
        }

        let legacyApi =
        navigator.getUserMedia ||
        navigator.webkitGetUserMedia ||
        navigator.mozGetUserMedia ||
        navigator.msGetUserMedia;

        if (legacyApi) {
            return new Promise(function (resolve, reject) {
                legacyApi.bind(window.navigator)(constraints, resolve, reject);
            });
        } else {
            alert("user api not supported");
        }
    }
    
    isRecording.textContent = "Recording...";
    
    let audioChunks = [];
    let rec;
    
    function handlerFunction(stream) {
        rec = new MediaRecorder(stream);
        rec.start();
        rec.ondataavailable = (e) => {
            audioChunks.push(e.data);
            if (rec.state == "inactive") {
                let blob = new Blob(audioChunks,{ 'type' : 'audio/wav; codecs=0' });
                document.getElementById("audioElement").src = URL.createObjectURL(blob);
                // sendAudioToServer(blob);
                // downloadAudio(blob, "test"); // Panggil fungsi downloadAudio dengan blob dan nama file yang diinginkan
            }
        };
    }
    
    function startusingBrowserMicrophone(boolean) {
        getUserMedia({ audio: boolean }).then((stream) => {
            handlerFunction(stream);
        });
    }
    
    startusingBrowserMicrophone(true);

    // Stoping handler
    document.getElementById("stopRecording").addEventListener("click", (e) => {
        rec.stop();
        isRecording.textContent = "Click play button to start listening";
    });
}

// Fungsi untuk mendownload file audio
function downloadAudio(blob, filename) {
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;
    link.download = filename;
    link.style.display = "none";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}

function sendAudioToServer(blob) {
    let formData = new FormData();
    formData.append("audio", blob, "suara.wav");
  
    fetch("http://localhost:5000/microphone", {
      method: "POST",
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => {
        alert("Sukses")
    })
      .catch((error) => {
        alert(error);
      });
}