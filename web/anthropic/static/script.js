document.addEventListener('DOMContentLoaded', () => {
    const promptTextarea = document.getElementById('prompt');
    const submitButton = document.getElementById('submit');
    const responseDiv = document.getElementById('response');

    submitButton.addEventListener('click', async () => {
        const prompt = promptTextarea.value.trim();
        if (!prompt) {
            alert('Please enter a prompt');
            return;
        }

        submitButton.disabled = true;
        responseDiv.textContent = 'Loading...';

        try {
            const response = await fetch('/api/query', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ prompt }),
            });

            const data = await response.json();
            
            if (data.error) {
                responseDiv.textContent = `Error: ${data.error}`;
            } else {
                responseDiv.textContent = data.result;
            }
        } catch (error) {
            responseDiv.textContent = `Error: ${error.message}`;
        } finally {
            submitButton.disabled = false;
        }
    });
});
