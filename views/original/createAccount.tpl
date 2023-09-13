<style>
body {
  background: #007bff;
  background: linear-gradient(to right, #0062E6, #33AEFF);
}

html, body {
  height: 100%;
  margin: 0;
}

.card-img-left {
  width: 45%;
  /* Link to your background image using in the property below! */
  background: scroll center url('/static/img/1595038166748.jpg');
  background-size: 80%;
background-position: center;
background-repeat: no-repeat;
}

.btn-login {
  font-size: 0.9rem;
  letter-spacing: 0.05rem;
  padding: 0.75rem 1rem;
}

.btn-google {
  color: white !important;
  background-color: #ea4335;
}

.btn-facebook {
  color: white !important;
  background-color: #3b5998;
}
.main{
    display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

</style>
{{ template "base.tpl" . }}
{{define "content"}}
    <!-- <h1>創建帳號</h1>
    <form action="" method="post">
    <label for="name">Name:<input type="text" id="name" name="name"></label>
    <label for="account">User:<input type="text" id="account" name="account"></label>
    <label for="password">Password:<input type="password" id="password" name="password"></label>
    <input type="submit" value="創建">
</form> -->
<div class="container">
    <div class="main row">
      <div class="col-lg-10 col-xl-9 mx-auto">
        <div class="card flex-row my-5 border-0 shadow rounded-3 overflow-hidden">
          <div class="card-img-left d-none d-md-flex">
            <!-- Background image for card set in CSS! -->
          </div>
          <div class="card-body p-4 p-sm-5">
            <h5 class="card-title text-center mb-5 fw-light fs-5">Register</h5>
            <form action="" method="post">

              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="name" name="name" placeholder="myusername" required autofocus>
                <label for="name">Username</label>
              </div>

              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="account" name="account" placeholder="costcogogo">
                <label for="account">Account</label>
              </div>


              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="password" name="password" placeholder="Password">
                <label for="password">Password</label>
              </div>

              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="floatingPasswordConfirm" placeholder="Confirm Password">
                <label for="floatingPasswordConfirm">Confirm Password</label>
              </div>

              <!-- Error message for password mismatch -->
              <div id="passwordMismatchError" class="alert alert-danger d-none mb-3">
                密碼和確認密碼不相符，請重新填寫。 <!-- 使用Go模板語言輸出錯誤訊息 -->
              </div>

              <div class="d-grid mb-2">
                <!-- 在創建按鈕上添加 onclick 事件，調用 validatePassword 函式進行密碼確認檢查 -->
                <button class="btn btn-lg btn-primary btn-login fw-bold text-uppercase" type="submit" onclick="validatePasswordAndSubmit(event)">Register</button>
              </div>

              <a class="d-block text-center mt-2 small" href="/login">已經有帳號? 登入</a>

              <hr class="my-4">

              <div class="d-grid mb-2">
                <button class="btn btn-lg btn-google btn-login fw-bold text-uppercase" type="submit">
                  <i class="fab fa-google me-2"></i> 使用 Google 帳號註冊
                </button>
              </div>

            </form>
          </div>
        </div>
      </div>
    </div>
  </div>

<script>
  // 密碼確認函式
  function validatePasswordAndSubmit(event) {
    // 獲取輸入的密碼和確認密碼
    const password = document.getElementById('password').value;
    const passwordConfirm = document.getElementById('floatingPasswordConfirm').value;

    // 檢查密碼是否相符
    if (password !== passwordConfirm) {
      // 如果密碼不相符，彈出錯誤提示框並清除錯誤訊息
    //   alert("密碼和確認密碼不相符，請重新填寫。");
      document.getElementById('passwordMismatchError').classList.remove('d-none');
      event.preventDefault(); // 阻止表單提交
    } else {
      // 如果密碼相符，提交表單
      document.getElementById('passwordMismatchError').classList.add('d-none');
      document.getElementsByTagName('form')[0].submit(); // 提交表單
      alert("創建成功");
      document.getElementById('name').value = ""; // 清空輸入的資料
      document.getElementById('account').value = "";
      document.getElementById('password').value = "";
      document.getElementById('floatingPasswordConfirm').value = "";
    }
  }
</script>
{{end}}