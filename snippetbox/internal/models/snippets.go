package models

import (
	"errors"
	"database/sql"
	"time"
)

type Snippet struct{
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct{
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int)(int,error){
	  smt := `INSERT INTO snippets (title, content, created, expires)
             VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY));`
      result, err := m.DB.Exec(smt,title,content,expires)
	  if err != nil {
        return 0,err
	  }
	  
	  id, err := result.LastInsertId()
	  if err != nil {
		return 0, err
	  }

      return int(id), nil

}
func (m *SnippetModel) Get(id int)(Snippet,error){
	smt := `Select * from snippets
	 where expires > UTC_TIMESTAMP() and id = ?`

	 row := m.DB.QueryRow(smt,id)

	 var s Snippet
     //escanejarà i copiarà els valors de les columnes a les referències de memoria que li hem passat.
	 err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	 if err != nil{
		if errors.Is(err,sql.ErrNoRows){
			return Snippet{}, ErrNoRecord
		}else{
			return Snippet{}, err
		}
	 }


	 return s,nil
}
func (m *SnippetModel) Latest()([]Snippet,error){
   
	smt := `Select * from snippets
	 where expires > UTC_TIMESTAMP()
	  ORDER BY id DESC limit 10
	`
	rows, err := m.DB.Query(smt)
	if err != nil {
		return nil,err
	}
     //this is critical
	defer rows.Close()


	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

        if err != nil {
			return nil,err
		}

		snippets = append(snippets,s)
	}

	if err=rows.Err(); err != nil {
		return nil,err
	}


	return snippets,nil
}

