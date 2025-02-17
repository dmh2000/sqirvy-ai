document.addEventListener('DOMContentLoaded', function() {
    // Get references to DOM elements
    const modelSelects = document.querySelectorAll('.model-select');
    const promptInput = document.getElementById('prompt');
    const submitButton = document.getElementById('submit');
    const results = document.querySelectorAll('.result');

    // Fetch available models when page loads
    fetch('http://localhost:8080/models')
        .then(response => response.json())
        .then(data => {
            // Populate model select dropdowns
            modelSelects.forEach((select, index) => {
                const providerName = select.closest('.result-box').querySelector('.provider-name');
                
                data.models.forEach(model => {
                    const option = document.createElement('option');
                    option.value = model.name;
                    option.textContent = `${model.name} (${model.provider})`;
                    option.dataset.provider = model.provider;
                    select.appendChild(option);
                });

                // Set initial provider name
                const selectedOption = select.options[0];
                if (selectedOption) {
                    providerName.textContent = selectedOption.dataset.provider;
                }

                // Update provider name when selection changes
                select.addEventListener('change', function() {
                    const selectedOption = this.options[this.selectedIndex];
                    providerName.textContent = selectedOption.dataset.provider;
                });
            });
        })
        .catch(error => console.error('Error fetching models:', error));

    // Handle form submission
    submitButton.addEventListener('click', async function() {
        const prompt = promptInput.value.trim();
        if (!prompt) {
            alert('Please enter a prompt');
            return;
        }

        // Clear previous results
        results.forEach(result => result.value = 'Loading...');

        // Make requests for each selected model
        for (let i = 0; i < modelSelects.length; i++) {
            const model = modelSelects[i].value;
            const resultArea = results[i];

            try {
                const response = await fetch('http://localhost:8080/query', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        model: model,
                        prompt: prompt,
                        temperature: 50
                    }),
                });

                const data = await response.json();
                resultArea.value = data.result;
            } catch (error) {
                resultArea.value = 'Error: ' + error.message;
            }
        }
    });
});
