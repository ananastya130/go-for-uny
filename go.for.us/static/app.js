document.getElementById('btn').onclick = doSearch;
document.getElementById('q').onkeydown = function(e){ if(e.key==='Enter') doSearch(); }

function doSearch(){
  const q = document.getElementById('q').value.trim();
  if(!q) { alert('Введите запрос'); return; }
  fetch('/api/search?q=' + encodeURIComponent(q))
    .then(r => r.json())
    .then(data => {
      const out = document.getElementById('results');
      if(data.error){ out.innerHTML = '<div class="err">'+data.error+'</div>'; return; }
      if(data.length === 0){ out.innerHTML = '<div class="empty">Ничего не найдено</div>'; return; }
      let html = '<table><tr><th>#</th><th>Группа</th><th>Лидер</th><th>Жанр</th><th>Альбомов</th></tr>';
      data.forEach((b, i) => {
        html += `<tr><td>${i+1}</td><td>${escapeHtml(b.name)}</td><td>${escapeHtml(b.leader)}</td><td>${escapeHtml(b.genre)}</td><td>${b.album_count}</td></tr>`;
      });
      html += '</table>';
      out.innerHTML = html;
    })
    .catch(err => {
      document.getElementById('results').innerText = 'Ошибка: ' + err;
    });
}

function escapeHtml(s){ return s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;'); }
