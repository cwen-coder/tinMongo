<div>
  <ul class="am-nav am-nav-tabs">
    <li  {{if equal .active "home" }}class="am-active" {{end}}><a href="/db/home">统计</a></li>
    <li {{if equal .active "newCollection" }}class="am-active" {{end}}><a href="/db/newCollection">创建集合</a></li>
    <li><a href="/command">命令</a></li>
    <li><a href="/execute">执行代码</a></li>
    <li {{if equal .active "dbTransfer" }}class="am-active" {{end}}><a href="/db/dbTransfer">克隆</a></li>
    <li {{if equal .active "dbExport" }}class="am-active" {{end}}><a href="/db/dbExport">导出</a></li>
    <li {{if equal .active "dbImport" }}class="am-active" {{end}} ><a href="/db/dbImport">导入</a></li>
    <li {{if equal .active "dbUsers" }}class="am-active" {{end}} ><a href="/db/dbUsers">用户</a></li>
    <li><a href="#">操作</a></li>
    <!-- <li class="am-dropdown" data-am-dropdown>
      <a class="am-dropdown-toggle" data-am-dropdown-toggle href="javascript:;">
        菜单 <span class="am-icon-caret-down"></span>
      </a>
      <ul class="am-dropdown-content">
        ...
      </ul>
    </li> -->
  </ul>
</div>