async function generatePasswords() {
    const length = document.getElementById('length').value;
    const count = document.getElementById('count').value;
    const symbols = document.getElementById('symbols').checked;

    const response = await fetch(`/generate?length=${length}&count=${count}&symbols=${symbols}`);
    const data = await response.json();

    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = '';

    data.passwords.forEach(password => {
        const passwordGroup = document.createElement('div');
        passwordGroup.className = 'password-group';
        
        const rawPassword = document.createElement('div');
        rawPassword.className = 'password raw';
        rawPassword.textContent = password.raw;
        
        const formattedPassword = document.createElement('div');
        formattedPassword.className = 'password formatted';
        formattedPassword.textContent = password.formatted;
        
        passwordGroup.appendChild(rawPassword);
        passwordGroup.appendChild(formattedPassword);
        resultsDiv.appendChild(passwordGroup);
    });
}
