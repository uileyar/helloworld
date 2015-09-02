package main

import "github.com/astaxie/beego"

/*
func (r *Reader) ReadForm(maxMemory int64) (f *Form, err error) {
	form := &Form{make(map[string][]string), make(map[string][]*FileHeader)}
	defer func() {
		if err != nil {
			form.RemoveAll()
		}
	}()

	maxValueBytes := int64(10 << 20) // 10 MB is a lot of text.
	for {
		p, err := r.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		name := p.FormName()
		if name == "" {
			continue
		}
		filename := p.FileName()

		var b bytes.Buffer

		if filename == "" {
			// value, store as string in memory
			n, err := io.CopyN(&b, p, maxValueBytes)
			if err != nil && err != io.EOF {
				return nil, err
			}
			maxValueBytes -= n
			if maxValueBytes == 0 {
				return nil, errors.New("multipart: message too large")
			}
			form.Value[name] = append(form.Value[name], b.String())
			continue
		}

		// file, store in memory or on disk
		fh := &FileHeader{
			Filename: filename,
			Header:   p.Header,
		}
		n, err := io.CopyN(&b, p, maxMemory+1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n > maxMemory {
			// too big, write to disk and flush buffer
			file, err := ioutil.TempFile("", "multipart-")
			if err != nil {
				return nil, err
			}
			defer file.Close()
			_, err = io.Copy(file, io.MultiReader(&b, p))
			if err != nil {
				os.Remove(file.Name())
				return nil, err
			}
			fh.tmpfile = file.Name()
		} else {
			fh.content = b.Bytes()
			maxMemory -= n
		}
		form.File[name] = append(form.File[name], fh)
	}

	return form, nil
}
*/
func main() {
	beego.Run()
}
