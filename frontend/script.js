async function createRoom() {
    const roomName = document.getElementById('roomName').value;
    const response = await fetch('http://localhost:8080/api/rooms', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: roomName }),
    });
    const data = await response.json();
    alert(`Room created with ID: ${data.roomID}`);
}
async function checkDeviceStatus() {
    try {
        const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
        const videoTracks = stream.getVideoTracks();
        const audioTracks = stream.getAudioTracks();

        if (videoTracks.length > 0 && audioTracks.length > 0) {
            alert("Camera and microphone are ready!");
            return true;
        } else {
            alert("Camera or microphone is not available.");
            return false;
        }
    } catch (error) {
        alert("Error accessing camera/microphone: " + error.message);
        return false;
    }
}

async function joinRoom() {
    const isDeviceReady = await checkDeviceStatus();
    if (!isDeviceReady) return;

    const roomID = document.getElementById('roomID').value;
    const userName = document.getElementById('userName').value;

    const response = await fetch(`http://localhost:8080/api/rooms/${roomID}/join`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ userName }),
    });
    const data = await response.json();
    alert(`Joined room with token: ${data.token}`);

    const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
    document.getElementById('localVideo').srcObject = stream;

    updateMembers([userName]);
}

async function joinRoom() {
    const roomID = document.getElementById('roomID').value;
    const userName = document.getElementById('userName').value;
    const response = await fetch(`http://localhost:8080/api/rooms/${roomID}/join`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ userName }),
    });
    const data = await response.json();
    alert(`Joined room with token: ${data.token}`);

    // Kết nối camera/micro
    const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
    document.getElementById('localVideo').srcObject = stream;

    // Hiển thị danh sách thành viên
    updateMembers([userName]);
}

function updateMembers(members) {
    const membersDiv = document.getElementById('members');
    membersDiv.innerHTML = `<h3>Members:</h3><ul>${members.map(m => `<li>${m}</li>`).join('')}</ul>`;
}