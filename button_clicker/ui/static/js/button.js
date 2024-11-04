let lastClickTime = 0; 
const clickLimit = 0.034; // Разрешенные клики в секунду 
let clickCount = 0; 
let timer; 

function updateTimer(seconds) { 
    const timerDisplay = document.getElementById("timer"); 
    if (seconds > 0) { 
        timerDisplay.textContent = seconds; 
    } else { 
        clearInterval(timer); 
        timerDisplay.textContent = "Готово!"; 
    } 
} 

function handleClick() { 
    const now = Date.now(); 
    if (now - lastClickTime >= 1000 / clickLimit) { 
        lastClickTime = now; 
        clickCount++; 

        // Создание POST-запроса 
        fetch('/submit', { 
            method: 'POST', 
            headers: { 
                'Content-Type': 'application/json', 
            }, 
            body: JSON.stringify({ message: "Button clicked!" }) 
        }) 
        .then(response => response.json()) 
        .then(data => { 
            if (data.success) { 
                document.getElementById("clickCount").textContent = data.count; 
                if (data.flag) { 
                    document.getElementById("flag").textContent = data.flag; 
                } 

                // Установка таймера на 30 секунд 
                updateTimer(30); 
                let seconds = 30; 
                timer = setInterval(() => { 
                    seconds--; 
                    updateTimer(seconds); 
                }, 1000);

                // Если сервер возвращает новый URL, перенаправляем пользователя
            }
            if (!data.success) { 
                window.location.href = "/"; 
            }
        }) 
        .catch(error => console.error('Error:', error)); 
    } else { 
        console.log('Слишком много нажатий!'); 
        document.getElementById("timer").textContent = "Подождите..."; 
    } 
}