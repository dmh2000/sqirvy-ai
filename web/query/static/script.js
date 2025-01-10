document.addEventListener('DOMContentLoaded', () => {
    const providerSelect = document.getElementById('provider');
    const promptTextarea = document.getElementById('prompt');
    const submitButton = document.getElementById('submit');
    const responseDiv = document.getElementById('response');

    submitButton.addEventListener('click', async () => {
        const provider = providerSelect.value;
        const prompt = promptTextarea.value.trim();
        
        if (!prompt) {
            alert('Please enter a prompt');
            return;
        }

        submitButton.disabled = true;
        responseDiv.textContent = 'Loading...';

        try {
            const response = await fetch(`/api/${provider}?prompt=${encodeURIComponent(prompt)}`);
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
