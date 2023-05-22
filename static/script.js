document.addEventListener('DOMContentLoaded', function() {
    var form = document.getElementById('loginForm');
    form.addEventListener('submit', function(event) {
        event.preventDefault();

        var username = document.getElementById('username').value;
        var password = document.getElementById('password').value;

        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json' // 修改为 'application/json'
            },
            body: JSON.stringify({ // 使用 JSON.stringify 将数据转换为 JSON 字符串
                username: username,
                password: password
            })
        })
            .then(function(response) {
                return response.json();
            })
            .then(function(data) {
                if (data.error) {
                    document.getElementById('error').textContent = data.error;
                } else {
                    document.getElementById('error').textContent = '';
                    var token = data.token;
                    loginWithToken(token);
                }
            })
            .catch(function(error) {
                console.error('请求错误:', error);
            });
    });

    function loginWithToken(token) {
        fetch('/auth/protected', {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + token
            }
        })
            .then(function(response) {
                if (response.status === 200) {
                    return response.json();
                } else {
                    throw new Error('Request failed with status: ' + response.status);
                }
            })
            .then(function(data) {
                console.log(data.message);
            })
            .catch(function(error) {
                if (error.message === 'Unauthorized') {
                    console.error('请求未经授权');
                } else {
                    console.error('请求错误:', error.message);
                }
            });
    }
});
