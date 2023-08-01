<style>
/* 請將剛才提供的 CSS 放在這裡 */

.login {
    min-height: 100vh;
}

.bg-image {
    background-image: url('/static/img/1595038166748.jpg');
    background-size: 60%;
    background-position: center;
    background-repeat: no-repeat;
    background-color: #FAF0E6;
}



.login-heading {
    font-weight: 300;
}

.btn-login {
    font-size: 0.9rem;
    letter-spacing: 0.05rem;
    padding: 0.75rem 1rem;
}
</style>
{{ template "base.tpl" . }}
{{define "content"}}
    <!-- <h1>Welcome to the Home Page</h1>
    <form action="" method="post">
    <label for="account">User:<input type="text" id="account" name="account"></label>
    <label for="password">Password:<input type="password" id="password" name="password"></label>
    <input type="submit" value="登入">
</form>
<a class="btn btn-primary" href="/create" role="button">創建帳號</a> -->

<!-- Sign In Form -->

<div class="container-fluid ps-md-0">
    <div class=" row g-0">
      <div class="d-none d-md-flex col-md-4 col-lg-6 bg-image"></div>
      <div class="col-md-8 col-lg-6">
        <div class="login d-flex align-items-center py-5">
          <div class="container">
            <div class="row">
              <div class="col-md-9 col-lg-8 mx-auto">
                <h3 class="login-heading mb-4">Welcome back!</h3>
                    <form method="post" action="">
                        <div class="form-floating mb-3">
                        <input type="text" class="form-control" id="account" name="account" placeholder="costcogogo">
                        <label for="account">Account</label>
                        </div>
                        <div class="form-floating mb-3">
                        <input type="password" class="form-control" id="password" name="password" placeholder="Password">
                        <label for="password">Password</label>
                        </div>
                    
                        <div class="form-check mb-3">
                        <input class="form-check-input" type="checkbox" value="" id="rememberPasswordCheck">
                        <label class="form-check-label" for="rememberPasswordCheck">
                            Remember password
                        </label>
                        </div>
                    
                        <div class="d-grid">
                        <input type="submit" class="btn btn-lg btn-primary btn-login text-uppercase fw-bold mb-2" value="Sign in">
                        <div class="text-center">
                            <a class="small" href="/create">Create an Account</a>
                        </div>
                        </div>
                    </form>
                </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>


  <script type="text/javascript">
    // 頁面載入時執行
    document.addEventListener('DOMContentLoaded', function() {
      // 檢查 LocalStorage 是否有儲存帳號資訊
      const account = localStorage.getItem('account');
      if (account) {
        // 填入表單中的帳號欄位
        document.getElementById('account').value = account;
        // 勾選 Remember password 的勾選框
        document.getElementById('rememberPasswordCheck').checked = true;
      }
    });
  
    // 表單提交事件處理函式
    document.getElementById('loginForm').addEventListener('submit', function(event) {
  
      const account = document.getElementById('account').value;
      const rememberPasswordCheck = document.getElementById('rememberPasswordCheck').checked;
  
      if (rememberPasswordCheck) {
        // 如果 Remember password 被勾選，將帳號資訊儲存到 LocalStorage 中
        localStorage.setItem('account', account);
      } else {
        // 如果 Remember password 未勾選，從 LocalStorage 中移除帳號資訊
        localStorage.removeItem('account');
      }
      document.getElementById('password').value = '';
    });

    
  </script> 
  
  
{{end}}