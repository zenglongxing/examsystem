package models

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) GetRandomSingleChoice(qtype string, count int) ([]SingleChoice, error) {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	ss := []SingleChoice{}
	err := t.Find(bson.M{"type": qtype}).All(&ss)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	c := len(ss)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	results := []SingleChoice{}
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(ss))

		results = append(results, ss[rn])
		ss = append(ss[:rn], ss[rn+1:]...)
	}

	return results, err
}

func (this *DBManager) GetRandomMultipleChoice(qtype string, count int) ([]MultipleChoice, error) {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	mc := []MultipleChoice{}
	err := t.Find(bson.M{"type": qtype}).All(&mc)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	c := len(mc)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	results := []MultipleChoice{}
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(mc))

		results = append(results, mc[rn])
		mc = append(mc[:rn], mc[rn+1:]...)
	}

	return results, err
}

func (this *DBManager) GetRandomTrueFalse(qtype string, count int) ([]TrueFalse, error) {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	tf := []TrueFalse{}
	err := t.Find(bson.M{"type": qtype}).All(&tf)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	c := len(tf)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	results := []TrueFalse{}
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(tf))

		results = append(results, tf[rn])
		tf = append(tf[:rn], tf[rn+1:]...)
	}

	return results, err
}

func (this *DBManager) GetSingleChoiceByDiscription(discription string) ([]SingleChoice, error) {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	ss := []SingleChoice{}
	err := t.Find(bson.M{"discription": discription}).All(&ss)

	return ss, err
}

func (this *DBManager) AddSingleChoice(s *SingleChoice) error {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	scs, err := this.GetSingleChoiceByDiscription(s.Discription)
	if err != nil {
		return err
	}

	for _, v := range scs {
		if v.Type == s.Type &&
			v.Discription == s.Discription &&
			v.A == s.A &&
			v.B == s.B &&
			v.C == s.C &&
			v.D == s.D {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

func (this *DBManager) GetMultipleChoiceByDiscription(discription string) ([]MultipleChoice, error) {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	ms := []MultipleChoice{}
	err := t.Find(bson.M{"discription": discription}).All(&ms)

	return ms, err
}

func (this *DBManager) AddMultipleChoice(m *MultipleChoice) error {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	mcs, err := this.GetMultipleChoiceByDiscription(m.Discription)
	if err != nil {
		return err
	}

	for _, v := range mcs {
		if v.Type == m.Type &&
			v.Discription == m.Discription &&
			v.A == m.A &&
			v.B == m.B &&
			v.C == m.C &&
			v.D == m.D &&
			v.E == m.E &&
			v.F == m.F {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(m)
	if err != nil {
		return err
	}

	return nil
}

func (this *DBManager) GetTrueFalseByDiscription(discription string) ([]TrueFalse, error) {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	ts := []TrueFalse{}
	err := t.Find(bson.M{"discription": discription}).All(&ts)

	return ts, err
}

func (this *DBManager) AddTrueFalse(f *TrueFalse) error {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	tfs, err := this.GetTrueFalseByDiscription(f.Discription)
	if err != nil {
		return err
	}

	for _, v := range tfs {
		if v.Type == f.Type && v.Discription == f.Discription {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(f)
	if err != nil {
		return err
	}

	return nil
}

func (this *DBManager) GetSingleChoiceSummary() (map[string]int, error) {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	scs := []SingleChoice{}
	err := t.Find(nil).All(&scs)
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	for _, sc := range scs {
		if v, ok := results[sc.Type]; ok {
			results[sc.Type] = v + 1
		} else {
			results[sc.Type] = 1
		}
	}

	return results, err
}

func (this *DBManager) GetMultipleChoiceSummary() (map[string]int, error) {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	mcs := []MultipleChoice{}
	err := t.Find(nil).All(&mcs)
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	for _, mc := range mcs {
		if v, ok := results[mc.Type]; ok {
			results[mc.Type] = v + 1
		} else {
			results[mc.Type] = 1
		}
	}

	return results, err
}

func (this *DBManager) GetTrueFalseSummary() (map[string]int, error) {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	tfs := []TrueFalse{}
	err := t.Find(nil).All(&tfs)
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	for _, tf := range tfs {
		if v, ok := results[tf.Type]; ok {
			results[tf.Type] = v + 1
		} else {
			results[tf.Type] = 1
		}
	}

	return results, err
}
