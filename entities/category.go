//sesuaikan dengan tabel di db

package entities

import "time"

type Category struct {
    Id        int
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}