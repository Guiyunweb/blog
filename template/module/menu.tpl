<nav class="navbar">
    <div class="navbar-content">
        <!-- logo -->
        <div class="navbar-logo">
            <a href="/">{{.Title}}</a>
        </div>
        <!-- link -->
        <div class="navbar-link">
            <div class="navbar-btn">
                <div></div>
                <div></div>
                <div></div>
            </div>
            <ul class="navbar-list">
                {{range $index, $value := .Menu}}
                    <li class="navbar-list-item">
                        <a href="{{$value.Url}}" target="{{$value.Target}}">{{$value.Name}}</a>
                    </li>
                {{end}}
            </ul>
        </div>
    </div>
</nav>