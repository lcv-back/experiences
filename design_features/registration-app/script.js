document.getElementById('registerForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:3000/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, email, password }),
        });

        const result = await response.json();
        document.getElementById('message').textContent = result.message;

        if (response.ok) {
            document.getElementById('message').style.color = 'green';
        }
    } catch (error) {
        document.getElementById('message').textContent = 'Có lỗi xảy ra, vui lòng thử lại!';
    }
});