{{define "mine"}}
        <div class="tit01">
            <h3 class="tit">关注我</h3>
            <div class="gzwm">
                <ul>
                    <li><a class="email" href="tel:{{.mine.Phone}}" target="_blank">我的电话</a></li>
                    <li><a class="qq" href="mailto:{{.mine.Email}}" target="_blank">我的邮箱</a></li>
                    <li><a class="tel" href="http://wpa.qq.com/msgrd?v=3&uin={{.mine.QQ}}&site=qq&menu=yes"
                               target="_blank" title="如果QQ未被唤起请按照提示添加我的QQ">我的QQ</a></li>
                    <li><a class="prize" href="/error" title="个人奖项">个人奖项</a>
                    </li>
                </ul>
                </div>
        </div>
{{end}}